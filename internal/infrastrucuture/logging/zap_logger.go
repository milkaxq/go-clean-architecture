package logging

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewLogger() *ZapLogger {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	return &ZapLogger{logger: logger}
}

func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {
	zapFields := convertZapFields(fields)
	l.logger.Info(msg, zapFields...)
}

func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {
	zapFields := convertZapFields(fields)
	l.logger.Error(msg, zapFields...)
}

func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {
	zapFields := convertZapFields(fields)
	l.logger.Debug(msg, zapFields...)
}

func (l *ZapLogger) Warning(msg string, fields map[string]interface{}) {
	zapFields := convertZapFields(fields)
	l.logger.Warn(msg, zapFields...)
}

func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {
	zapFields := convertZapFields(fields)
	l.logger.Fatal(msg, zapFields...)
}

func convertZapFields(fields map[string]interface{}) []zap.Field {
	var zapField []zap.Field

	for k, v := range fields {
		zapField = append(zapField, zap.Any(k, v))
	}

	return zapField
}
