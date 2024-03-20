package context

import (
	"context"

	"github.com/edward-/four-in-a-row-game/pkg/logger"
)

func LoggerFromCtx(ctx context.Context) logger.Logger {
	return ctx.Value(loggerKey).(logger.Logger)
}

func SetLogger(ctx context.Context, logger logger.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
