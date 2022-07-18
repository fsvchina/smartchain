export
functionality.

Randomness

Currently, simulation uses a single seed (integer) as a source for a PRNG by
which all random operations are executed from. Any call to the PRNG changes all
future operations as the internal state of the PRNG is modified. For example,
if a new message type is created and needs to be simulated, the new introduced
PRNG call will change all subsequent operations.

This may can often be problematic when testing fixes to simulation faults. One
current solution to this is to use a params file as mentioned above. In the future
the simulation suite is expected to support a series of PRNGs that can be used
uniquely per module and simulation component so that they will not effect each
others state execution outcome.

Usage

To execute a completely pseudo-random simulation:

 $ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
	-run=TestFullAppSimulation \
	-Enabled=true \
	-NumBlocks=100 \
	-BlockSize=200 \
	-Commit=true \
	-Seed=99 \
	-Period=5 \
	-v -timeout 24h

To execute simulation from a genesis file:

 $ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
 	-run=TestFullAppSimulation \
 	-Enabled=true \
 	-NumBlocks=100 \
 	-BlockSize=200 \
 	-Commit=true \
 	-Seed=99 \
 	-Period=5 \
	-Genesis=/path/to/genesis.json \
 	-v -timeout 24h

To execute simulation from a simulation params file:

 $ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
	-run=TestFullAppSimulation \
	-Enabled=true \
	-NumBlocks=100 \
	-BlockSize=200 \
	-Commit=true \
	-Seed=99 \
	-Period=5 \
	-Params=/path/to/params.json \
	-v -timeout 24h

To export the simulation params to a file at a given block height:

 $ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
 	-run=TestFullAppSimulation \
 	-Enabled=true \
 	-NumBlocks=100 \
 	-BlockSize=200 \
 	-Commit=true \
 	-Seed=99 \
 	-Period=5 \
	-ExportParamsPath=/path/to/params.json \
	-ExportParamsHeight=50 \
	 -v -timeout 24h


To export the simulation app state (i.e genesis) to a file:

 $ go test -mod=readonly github.com/cosmos/cosmos-sdk/simapp \
 	-run=TestFullAppSimulation \
 	-Enabled=true \
 	-NumBlocks=100 \
 	-BlockSize=200 \
 	-Commit=true \
 	-Seed=99 \
 	-Period=5 \
	-ExportStatePath=/path/to/genesis.json \
	 v -timeout 24h

Params

Params that are provided to simulation from a JSON file are used to used to set
both module parameters and simulation parameters. See sim_test.go for the full
set of parameters that can be provided.
*/
package simulation
