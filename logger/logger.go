package logger

import "log"

type ILogger interface {
	LogInfo(event interface{})
}

type Logger struct{}

func CreateLogger() ILogger {
	return &Logger{}
}

func (s *Logger) LogInfo(event interface{}) {
	l := log.Default()
	l.Println(event)
}
