

package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultPage    = 1
	DefaultLimit   = 30             
	TxMinHeightKey = "tx.minheight" 
	TxMaxHeightKey = "tx.maxheight" 
)



type ResponseWithHeight struct {
	Height int64           `json:"height"`
	Result json.RawMessage `json:"result"`
}


func NewResponseWithHeight(height int64, result json.RawMessage) ResponseWithHeight {
	return ResponseWithHeight{
		Height: height,
		Result: result,
	}
}



func ParseResponseWithHeight(cdc *codec.LegacyAmino, bz []byte) ([]byte, error) {
	r := ResponseWithHeight{}
	if err := cdc.UnmarshalJSON(bz, &r); err != nil {
		return nil, err
	}

	return r.Result, nil
}


type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate"`
}



type BaseReq struct {
	From          string       `json:"from"`
	Memo          string       `json:"memo"`
	ChainID       string       `json:"chain_id"`
	AccountNumber uint64       `json:"account_number"`
	Sequence      uint64       `json:"sequence"`
	TimeoutHeight uint64       `json:"timeout_height"`
	Fees          sdk.Coins    `json:"fees"`
	GasPrices     sdk.DecCoins `json:"gas_prices"`
	Gas           string       `json:"gas"`
	GasAdjustment string       `json:"gas_adjustment"`
	Simulate      bool         `json:"simulate"`
}


func NewBaseReq(
	from, memo, chainID string, gas, gasAdjustment string, accNumber, seq uint64,
	fees sdk.Coins, gasPrices sdk.DecCoins, simulate bool,
) BaseReq {
	return BaseReq{
		From:          strings.TrimSpace(from),
		Memo:          strings.TrimSpace(memo),
		ChainID:       strings.TrimSpace(chainID),
		Fees:          fees,
		GasPrices:     gasPrices,
		Gas:           strings.TrimSpace(gas),
		GasAdjustment: strings.TrimSpace(gasAdjustment),
		AccountNumber: accNumber,
		Sequence:      seq,
		Simulate:      simulate,
	}
}


func (br BaseReq) Sanitize() BaseReq {
	return NewBaseReq(
		br.From, br.Memo, br.ChainID, br.Gas, br.GasAdjustment,
		br.AccountNumber, br.Sequence, br.Fees, br.GasPrices, br.Simulate,
	)
}




func (br BaseReq) ValidateBasic(w http.ResponseWriter) bool {
	if !br.Simulate {
		switch {
		case len(br.ChainID) == 0:
			WriteErrorResponse(w, http.StatusUnauthorized, "chain-id required but not specified")
			return false

		case !br.Fees.IsZero() && !br.GasPrices.IsZero():
			
			WriteErrorResponse(w, http.StatusBadRequest, "cannot provide both fees and gas prices")
			return false

		case !br.Fees.IsValid() && !br.GasPrices.IsValid():
			
			WriteErrorResponse(w, http.StatusPaymentRequired, "invalid fees or gas prices provided")
			return false
		}
	}

	if _, err := sdk.AccAddressFromBech32(br.From); err != nil || len(br.From) == 0 {
		WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("invalid from address: %s", br.From))
		return false
	}

	return true
}



func ReadRESTReq(w http.ResponseWriter, r *http.Request, cdc *codec.LegacyAmino, req interface{}) bool {
	body, err := ioutil.ReadAll(r.Body)
	if CheckBadRequestError(w, err) {
		return false
	}

	err = cdc.UnmarshalJSON(body, req)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to decode JSON payload: %s", err))
		return false
	}

	return true
}


type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}


func NewErrorResponse(code int, err string) ErrorResponse {
	return ErrorResponse{Code: code, Error: err}
}



func CheckError(w http.ResponseWriter, status int, err error) bool {
	if err != nil {
		WriteErrorResponse(w, status, err.Error())
		return true
	}

	return false
}



func CheckBadRequestError(w http.ResponseWriter, err error) bool {
	return CheckError(w, http.StatusBadRequest, err)
}



func CheckInternalServerError(w http.ResponseWriter, err error) bool {
	return CheckError(w, http.StatusInternalServerError, err)
}



func CheckNotFoundError(w http.ResponseWriter, err error) bool {
	return CheckError(w, http.StatusNotFound, err)
}



func WriteErrorResponse(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(legacy.Cdc.MustMarshalJSON(NewErrorResponse(0, err)))
}



func WriteSimulationResponse(w http.ResponseWriter, cdc *codec.LegacyAmino, gas uint64) {
	gasEst := GasEstimateResponse{GasEstimate: gas}

	resp, err := cdc.MarshalJSON(gasEst)
	if CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}


func ParseUint64OrReturnBadRequest(w http.ResponseWriter, s string) (n uint64, ok bool) {
	var err error

	n, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("'%s' is not a valid uint64", s))

		return n, false
	}

	return n, true
}



func ParseFloat64OrReturnBadRequest(w http.ResponseWriter, s string, defaultIfEmpty float64) (n float64, ok bool) {
	if len(s) == 0 {
		return defaultIfEmpty, true
	}

	n, err := strconv.ParseFloat(s, 64)
	if CheckBadRequestError(w, err) {
		return n, false
	}

	return n, true
}



func ParseQueryHeightOrReturnBadRequest(w http.ResponseWriter, clientCtx client.Context, r *http.Request) (client.Context, bool) {
	heightStr := r.FormValue("height")
	if heightStr != "" {
		height, err := strconv.ParseInt(heightStr, 10, 64)
		if CheckBadRequestError(w, err) {
			return clientCtx, false
		}

		if height < 0 {
			WriteErrorResponse(w, http.StatusBadRequest, "height must be equal or greater than zero")
			return clientCtx, false
		}

		if height > 0 {
			clientCtx = clientCtx.WithHeight(height)
		}
	} else {
		clientCtx = clientCtx.WithHeight(0)
	}

	return clientCtx, true
}



func PostProcessResponseBare(w http.ResponseWriter, ctx client.Context, body interface{}) {
	var (
		resp []byte
		err  error
	)

	switch b := body.(type) {
	case []byte:
		resp = b

	default:
		resp, err = ctx.LegacyAmino.MarshalJSON(body)
		if CheckInternalServerError(w, err) {
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}




func PostProcessResponse(w http.ResponseWriter, ctx client.Context, resp interface{}) {
	var (
		result []byte
		err    error
	)

	if ctx.Height < 0 {
		WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("negative height in response").Error())
		return
	}

	
	marshaler := ctx.LegacyAmino

	switch res := resp.(type) {
	case []byte:
		result = res

	default:
		result, err = marshaler.MarshalJSON(resp)
		if CheckInternalServerError(w, err) {
			return
		}
	}

	wrappedResp := NewResponseWithHeight(ctx.Height, result)

	output, err := marshaler.MarshalJSON(wrappedResp)
	if CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}




func ParseHTTPArgsWithLimit(r *http.Request, defaultLimit int) (tags []string, page, limit int, err error) {
	tags = make([]string, 0, len(r.Form))

	for key, values := range r.Form {
		if key == "page" || key == "limit" {
			continue
		}

		var value string
		value, err = url.QueryUnescape(values[0])

		if err != nil {
			return tags, page, limit, err
		}

		var tag string

		switch key {
		case types.TxHeightKey:
			tag = fmt.Sprintf("%s=%s", key, value)

		case TxMinHeightKey:
			tag = fmt.Sprintf("%s>=%s", types.TxHeightKey, value)

		case TxMaxHeightKey:
			tag = fmt.Sprintf("%s<=%s", types.TxHeightKey, value)

		default:
			tag = fmt.Sprintf("%s='%s'", key, value)
		}

		tags = append(tags, tag)
	}

	pageStr := r.FormValue("page")
	if pageStr == "" {
		page = DefaultPage
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return tags, page, limit, err
		} else if page <= 0 {
			return tags, page, limit, errors.New("page must greater than 0")
		}
	}

	limitStr := r.FormValue("limit")
	if limitStr == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return tags, page, limit, err
		} else if limit <= 0 {
			return tags, page, limit, errors.New("limit must greater than 0")
		}
	}

	return tags, page, limit, nil
}



func ParseHTTPArgs(r *http.Request) (tags []string, page, limit int, err error) {
	return ParseHTTPArgsWithLimit(r, DefaultLimit)
}



func ParseQueryParamBool(r *http.Request, param string) bool {
	if value, err := strconv.ParseBool(r.FormValue(param)); err == nil {
		return value
	}

	return false
}



func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url) 
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	return body, nil
}



func PostRequest(url string, contentType string, data []byte) ([]byte, error) {
	res, err := http.Post(url, contentType, bytes.NewBuffer(data)) 
	if err != nil {
		return nil, fmt.Errorf("error while sending post request: %w", err)
	}

	bz, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	return bz, nil
}
