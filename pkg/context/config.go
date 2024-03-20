package context

import (
	"context"

	"github.com/edward-/four-in-a-row-game/pkg/config"
)

func ConfigFromCtx(ctx context.Context) config.Config {
	return ctx.Value(configkey).(config.Config)
}

func SetConfig(ctx context.Context, cfg config.Config) context.Context {
	return context.WithValue(ctx, configkey, cfg)
}
