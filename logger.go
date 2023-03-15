package waterslog

import (
	"github.com/ThreeDotsLabs/watermill"
	"golang.org/x/exp/slog"
)

// forces the compiler to check if the logger implements the interface
var _ watermill.LoggerAdapter = (*WaterSLogger)(nil)

type WaterSLogger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *WaterSLogger {
	return &WaterSLogger{
		logger: logger,
	}
}

func (w *WaterSLogger) Error(msg string, err error, fields watermill.LogFields) {
	var kv []interface{}
	if fields != nil {
		kv = keyValFromFields(fields)
	}
	kv = append(kv, "error", err)
	w.logger.Error(msg, kv...)
}

func (w *WaterSLogger) Info(msg string, fields watermill.LogFields) {
	if fields != nil {
		kv := keyValFromFields(fields)
		w.logger.Info(msg, kv...)
	} else {
		w.logger.Info(msg)
	}
}

func (w *WaterSLogger) Debug(msg string, fields watermill.LogFields) {
	if fields != nil {
		kv := keyValFromFields(fields)
		w.logger.Debug(msg, kv...)
	} else {
		w.logger.Debug(msg)
	}
}

func (w *WaterSLogger) Trace(msg string, fields watermill.LogFields) {
	w.Debug(msg, fields)
}

func (w *WaterSLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	if fields == nil {
		return w
	}

	return &WaterSLogger{
		logger: w.logger.With(keyValFromFields(fields)...),
	}
}

func keyValFromFields(fields watermill.LogFields) []interface{} {
	if fields == nil {
		return nil
	}

	var keyVal []interface{}
	for k, v := range fields {
		keyVal = append(keyVal, k, v)
	}
	return keyVal
}
