package dmetering

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/streamingfast/dmetering", &zlog)
}
