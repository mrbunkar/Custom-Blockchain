package util

import "go.uber.org/zap"

type Logger interface {
	Log(string, string)
}

type ZapLogger struct {
	*zap.Logger
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{
		&zap.Logger{},
	}
}

func (zlog *ZapLogger) Log(level string, msg string) {

}
