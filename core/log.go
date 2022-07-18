package core

import (
	"fs.video/log"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	LmChainClient = log.RegisterModule("sccli", logrus.InfoLevel)
)


func BuildLog(funcName string, modules ...log.LogModule) *logrus.Entry {
	moduleName := ""
	for _, v := range modules {
		if moduleName != "" {
			moduleName += "/"
		}
		moduleName += string(v)
	}
	logEntry := log.Log.WithField("module", strings.ToLower(moduleName))
	if funcName != "" {
		logEntry = logEntry.WithField("method", strings.ToLower(funcName))
	}
	return logEntry
}
