package main

import (
	"github.com/xpro-development/xpro-utility/logger"
	"strings"
)

var logs logger.ILogger

func WriteTestLog(message string) {
	logs.LogInfo(strings.ToUpper(message))
}

func init() {
	logs = logger.CreateLogger()
	CreateNewUtility(logs)
}

func main() {
	WriteTestLog("hello from main method")
}
