package logging

import (
	"go.uber.org/zap"
)

// Logger used for the whole application.
var Logger *zap.Logger

func init() {
	Logger, _ = zap.NewProduction()
}
