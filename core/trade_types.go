package core

import (
	"fs.video/trerr"
)



var (
	TradeTypeTransfer       = RegisterTranserType("transfer", "轉帳", "transfer accounts")
	TradeTypeIbcTransferOut = RegisterTranserType("ibc-transfer-out", "IBC跨鏈轉出", "IBC Transfer Out")
	TradeTypeIbcTransferIn  = RegisterTranserType("ibc-transfer-in", "IBC跨鏈轉入", "IBC Transfer In")
)

var tradeTypeText = make(map[string]string)
var tradeTypeTextEn = make(map[string]string)


func RegisterTranserType(key, value, enValue string) TranserType {
	tradeTypeTextEn[key] = enValue
	tradeTypeText[key] = value
	return TranserType(key)
}


func GetTranserTypeConfig() map[string]string {
	if trerr.Language == "EN" {
		return tradeTypeTextEn
	} else {
		return tradeTypeText
	}
}

type TranserType string

func (this TranserType) GetValue() string {
	if text, ok := tradeTypeText[string(this[:])]; ok {
		return text
	} else {
		return ""
	}
}

func (this TranserType) GetKey() string {
	return string(this[:])
}
