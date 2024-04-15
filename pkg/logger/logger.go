package logger

import "go.uber.org/zap"

//there gonna be the logger configuration later

type Logger struct {
	zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}
