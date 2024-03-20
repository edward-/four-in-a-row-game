package context

import (
	"context"

	"gorm.io/gorm"
)

func DatabaseFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value(databaseKey).(*gorm.DB)
}

func SetDatabase(ctx context.Context, gormInstance *gorm.DB) context.Context {
	return context.WithValue(ctx, databaseKey, gormInstance)
}
