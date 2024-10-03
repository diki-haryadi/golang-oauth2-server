package cronJob

import "golang-oauth2-server/pkg/logger"

type CronLogger struct{}

func NewLogger() *CronLogger {
	return &CronLogger{}
}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Infow(msg, keysAndValues)
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Errorw(msg, keysAndValues)
}
