package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"reflect"
	"runtime"
	"strings"
)


func GetStructFuncName(object interface{}) string {
	structName := reflect.TypeOf(object).String()
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	arry := strings.Split(f.Name(), ".")
	funcName := arry[len(arry)-1]
	return structName + "." + funcName
}



func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	arry := strings.Split(f.Name(), ".")
	return arry[len(arry)-1]
}




func GetPackageFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	list := strings.Split(f.Name(), "/")
	return list[len(list)-1]
}


func GetEscrowAddress(sourcePort, sourceChannel string) sdk.AccAddress {
	return types.GetEscrowAddress(sourcePort, sourceChannel)
}
