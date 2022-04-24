package main

import "github.com/xpro-development/xpro-utility/logger"

type IUtility interface {
}

type Utility struct {
	logger logger.ILogger
}

func CreateNewUtility(logger logger.ILogger) IUtility {
	return &Utility{
		logger: logger,
	}
}
