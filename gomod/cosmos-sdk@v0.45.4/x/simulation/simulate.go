package simulation

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
)

const AverageBlockTime = 6 * time.Second


func initChain(
	r *rand.Rand,
	params Params,
	accounts []simulation.Account,
	app *baseapp.BaseApp,
	appStateFn simulation.AppStateFn,
	config simulation.Config,
	cdc codec.JSONCodec,
) (mockValidators, time.Time, []simulation.Account, string) {

	appState, accounts, chainID, genesisTimestamp := appStateFn(r, accounts, config)
	consensusParams := randomConsensusParams(r, appState, cdc)
	req := abci.RequestInitChain{
		AppStateBytes:   appState,
		ChainId:         chainID,
		ConsensusParams: consensusParams,
		Time:            genesisTimestamp,
	}
	res := app.InitChain(req)
	validators := newMockValidators(r, res.Validators, params)

	return validators, genesisTimestamp, accounts, chainID
}




func SimulateFromSeed(
	tb testing.TB,
	w io.Writer,
	app *baseapp.BaseApp,
	appStateFn simulation.AppStateFn,
	randAccFn simulation.RandomAccountFn,
	ops WeightedOperations,
	blockedAddrs map[string]bool,
	config simulation.Config,
	cdc codec.JSONCodec,
) (stopEarly bool, exportedParams Params, err error) {

	testingMode, _, b := getTestingMode(tb)

	fmt.Fprintf(w, "Starting SimulateFromSeed with randomness created with seed %d\n", int(config.Seed))
	r := rand.New(rand.NewSource(config.Seed))
	params := RandomParams(r)
	fmt.Fprintf(w, "Randomized simulation params: \n%s\n", mustMarshalJSONIndent(params))

	timeDiff := maxTimePerBlock - minTimePerBlock
	accs := randAccFn(r, params.NumKeys())
	eventStats := NewEventStats()



	validators, genesisTimestamp, accs, chainID := initChain(r, params, accs, app, appStateFn, config, cdc)
	if len(accs) == 0 {
		return true, params, fmt.Errorf("must have greater than zero genesis accounts")
	}

	config.ChainID = chainID

	fmt.Printf(
		"Starting the simulation from time %v (unixtime %v)\n",
		genesisTimestamp.UTC().Format(time.UnixDate), genesisTimestamp.Unix(),
	)


	var tmpAccs []simulation.Account

	for _, acc := range accs {
		if !blockedAddrs[acc.Address.String()] {
			tmpAccs = append(tmpAccs, acc)
		}
	}

	accs = tmpAccs
	nextValidators := validators

	header := tmproto.Header{
		ChainID:         config.ChainID,
		Height:          1,
		Time:            genesisTimestamp,
		ProposerAddress: validators.randomProposer(r),
	}
	opCount := 0


	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		receivedSignal := <-c
		fmt.Fprintf(w, "\nExiting early due to %s, on block %d, operation %d\n", receivedSignal, header.Height, opCount)
		err = fmt.Errorf("exited due to %s", receivedSignal)
		stopEarly = true
	}()

	var (
		pastTimes     []time.Time
		pastVoteInfos [][]abci.VoteInfo
	)

	request := RandomRequestBeginBlock(r, params,
		validators, pastTimes, pastVoteInfos, eventStats.Tally, header)


	operationQueue := NewOperationQueue()

	var timeOperationQueue []simulation.FutureOperation

	logWriter := NewLogWriter(testingMode)

	blockSimulator := createBlockSimulator(
		testingMode, tb, w, params, eventStats.Tally,
		ops, operationQueue, timeOperationQueue, logWriter, config)

	if !testingMode {
		b.ResetTimer()
	} else {

		defer func() {
			if r := recover(); r != nil {
				_, _ = fmt.Fprintf(w, "simulation halted due to panic on block %d\n", header.Height)
				logWriter.PrintLogs()
				panic(r)
			}
		}()
	}


	if config.ExportParamsPath != "" && config.ExportParamsHeight == 0 {
		exportedParams = params
	}


	for height := config.InitialBlockHeight; height < config.NumBlocks+config.InitialBlockHeight && !stopEarly; height++ {


		pastTimes = append(pastTimes, header.Time)
		pastVoteInfos = append(pastVoteInfos, request.LastCommitInfo.Votes)


		logWriter.AddEntry(BeginBlockEntry(int64(height)))
		app.BeginBlock(request)

		ctx := app.NewContext(false, header)


		numQueuedOpsRan := runQueuedOperations(
			operationQueue, int(header.Height), tb, r, app, ctx, accs, logWriter,
			eventStats.Tally, config.Lean, config.ChainID,
		)

		numQueuedTimeOpsRan := runQueuedTimeOperations(
			timeOperationQueue, int(header.Height), header.Time,
			tb, r, app, ctx, accs, logWriter, eventStats.Tally,
			config.Lean, config.ChainID,
		)


		operations := blockSimulator(r, app, ctx, accs, header)
		opCount += operations + numQueuedOpsRan + numQueuedTimeOpsRan

		res := app.EndBlock(abci.RequestEndBlock{})
		header.Height++
		header.Time = header.Time.Add(
			time.Duration(minTimePerBlock) * time.Second)
		header.Time = header.Time.Add(
			time.Duration(int64(r.Intn(int(timeDiff)))) * time.Second)
		header.ProposerAddress = validators.randomProposer(r)

		logWriter.AddEntry(EndBlockEntry(int64(height)))

		if config.Commit {
			app.Commit()
		}

		if header.ProposerAddress == nil {
			fmt.Fprintf(w, "\nSimulation stopped early as all validators have been unbonded; nobody left to propose a block!\n")
			stopEarly = true
			break
		}



		request = RandomRequestBeginBlock(r, params, validators, pastTimes, pastVoteInfos, eventStats.Tally, header)



		validators = nextValidators
		nextValidators = updateValidators(tb, r, params, validators, res.ValidatorUpdates, eventStats.Tally)


		if config.ExportParamsPath != "" && config.ExportParamsHeight == height {
			exportedParams = params
		}
	}

	if stopEarly {
		if config.ExportStatsPath != "" {
			fmt.Println("Exporting simulation statistics...")
			eventStats.ExportJSON(config.ExportStatsPath)
		} else {
			eventStats.Print(w)
		}

		return true, exportedParams, err
	}

	fmt.Fprintf(
		w,
		"\nSimulation complete; Final height (blocks): %d, final time (seconds): %v, operations ran: %d\n",
		header.Height, header.Time, opCount,
	)

	if config.ExportStatsPath != "" {
		fmt.Println("Exporting simulation statistics...")
		eventStats.ExportJSON(config.ExportStatsPath)
	} else {
		eventStats.Print(w)
	}

	return false, exportedParams, nil
}

type blockSimFn func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
	accounts []simulation.Account, header tmproto.Header) (opCount int)



func createBlockSimulator(testingMode bool, tb testing.TB, w io.Writer, params Params,
	event func(route, op, evResult string), ops WeightedOperations,
	operationQueue OperationQueue, timeOperationQueue []simulation.FutureOperation,
	logWriter LogWriter, config simulation.Config) blockSimFn {

	lastBlockSizeState := 0
	blocksize := 0
	selectOp := ops.getSelectOpFn()

	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, header tmproto.Header,
	) (opCount int) {
		_, _ = fmt.Fprintf(
			w, "\rSimulating... block %d/%d, operation %d/%d.",
			header.Height, config.NumBlocks, opCount, blocksize,
		)
		lastBlockSizeState, blocksize = getBlockSize(r, params, lastBlockSizeState, config.BlockSize)

		type opAndR struct {
			op   simulation.Operation
			rand *rand.Rand
		}

		opAndRz := make([]opAndR, 0, blocksize)



		for i := 0; i < blocksize; i++ {
			opAndRz = append(opAndRz, opAndR{
				op:   selectOp(r),
				rand: simulation.DeriveRand(r),
			})
		}

		for i := 0; i < blocksize; i++ {

			opAndR := opAndRz[i]
			op, r2 := opAndR.op, opAndR.rand
			opMsg, futureOps, err := op(r2, app, ctx, accounts, config.ChainID)
			opMsg.LogEvent(event)

			if !config.Lean || opMsg.OK {
				logWriter.AddEntry(MsgEntry(header.Height, int64(i), opMsg))
			}

			if err != nil {
				logWriter.PrintLogs()
				tb.Fatalf(`error on block  %d/%d, operation (%d/%d) from x/%s:
%v
Comment: %s`,
					header.Height, config.NumBlocks, opCount, blocksize, opMsg.Route, err, opMsg.Comment)
			}

			queueOperations(operationQueue, timeOperationQueue, futureOps)

			if testingMode && opCount%50 == 0 {
				fmt.Fprintf(w, "\rSimulating... block %d/%d, operation %d/%d. ",
					header.Height, config.NumBlocks, opCount, blocksize)
			}

			opCount++
		}

		return opCount
	}
}


func runQueuedOperations(queueOps map[int][]simulation.Operation,
	height int, tb testing.TB, r *rand.Rand, app *baseapp.BaseApp,
	ctx sdk.Context, accounts []simulation.Account, logWriter LogWriter,
	event func(route, op, evResult string), lean bool, chainID string) (numOpsRan int) {

	queuedOp, ok := queueOps[height]
	if !ok {
		return 0
	}

	numOpsRan = len(queuedOp)
	for i := 0; i < numOpsRan; i++ {




		opMsg, _, err := queuedOp[i](r, app, ctx, accounts, chainID)
		opMsg.LogEvent(event)

		if !lean || opMsg.OK {
			logWriter.AddEntry((QueuedMsgEntry(int64(height), opMsg)))
		}

		if err != nil {
			logWriter.PrintLogs()
			tb.FailNow()
		}
	}
	delete(queueOps, height)

	return numOpsRan
}

func runQueuedTimeOperations(queueOps []simulation.FutureOperation,
	height int, currentTime time.Time, tb testing.TB, r *rand.Rand,
	app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account,
	logWriter LogWriter, event func(route, op, evResult string),
	lean bool, chainID string) (numOpsRan int) {

	numOpsRan = 0
	for len(queueOps) > 0 && currentTime.After(queueOps[0].BlockTime) {




		opMsg, _, err := queueOps[0].Op(r, app, ctx, accounts, chainID)
		opMsg.LogEvent(event)

		if !lean || opMsg.OK {
			logWriter.AddEntry(QueuedMsgEntry(int64(height), opMsg))
		}

		if err != nil {
			logWriter.PrintLogs()
			tb.FailNow()
		}

		queueOps = queueOps[1:]
		numOpsRan++
	}

	return numOpsRan
}
