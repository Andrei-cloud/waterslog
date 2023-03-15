package waterslog_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/andrei-cloud/waterslog"
	"github.com/stretchr/testify/assert"

	"github.com/ThreeDotsLabs/watermill"
	"golang.org/x/exp/slog"
)

func NewTestLogger(out io.Writer) *slog.Logger {
	l := slog.New(slog.HandlerOptions{Level: slog.LevelDebug}.NewJSONHandler(out))
	return l
}

func TestWaterSLogger(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	logger := waterslog.New(NewTestLogger(buf))

	t.Run("INFO", func(t *testing.T) {
		buf.Reset()
		logger.Info("hello", nil)
		assert.Contains(t, buf.String(), `"level":"INFO"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
	})

	t.Run("INFO", func(t *testing.T) {
		buf.Reset()
		logger.Info("hello", watermill.LogFields{"foo": "bar"})
		assert.Contains(t, buf.String(), `"level":"INFO"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
		assert.Contains(t, buf.String(), `"foo":"bar"`)
	})

	t.Run("ERROR", func(t *testing.T) {
		buf.Reset()
		logger.Error("hello", errors.New("some error"), nil)
		assert.Contains(t, buf.String(), `"level":"ERROR"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
		assert.Contains(t, buf.String(), `"error":"some error"`)
	})

	t.Run("ERROR", func(t *testing.T) {
		buf.Reset()
		logger.Error("hello", errors.New("some error"), watermill.LogFields{"foo": "bar"})
		assert.Contains(t, buf.String(), `"level":"ERROR"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
		assert.Contains(t, buf.String(), `"foo":"bar"`)
		assert.Contains(t, buf.String(), `"error":"some error"`)
	})

	t.Run("DEBUG", func(t *testing.T) {
		buf.Reset()
		logger.Debug("hello", nil)
		assert.Contains(t, buf.String(), `"level":"DEBUG"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
	})

	t.Run("DEBUG", func(t *testing.T) {
		buf.Reset()
		logger.Debug("hello", watermill.LogFields{"foo": "bar"})
		assert.Contains(t, buf.String(), `"level":"DEBUG"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
		assert.Contains(t, buf.String(), `"foo":"bar"`)
	})
}

func TestWaterSLogger_With(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	logger := waterslog.New(NewTestLogger(buf)).With(watermill.LogFields{"tag": "always"})

	t.Run("INFO", func(t *testing.T) {
		buf.Reset()
		logger.Info("hello", nil)
		assert.Contains(t, buf.String(), `"level":"INFO"`)
		assert.Contains(t, buf.String(), `"tag":"always"`)
		assert.Contains(t, buf.String(), `"msg":"hello"`)
	})
}
