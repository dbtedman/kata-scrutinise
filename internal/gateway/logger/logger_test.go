package logger_test

import (
	"github.com/dbtedman/kata-scrutinise/internal/gateway/logger"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

func TestContextFields(t *testing.T) {
	t.Run("defines expected context fields", func(t *testing.T) {
		fields := logger.ContextFields()
		assert.Assert(t, is.Len(fields, 0))
	})
}
