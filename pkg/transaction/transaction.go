package transaction

import (
	"context"

	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"

	"gorm.io/gorm"
)

type TransactionFunc func(ctx context.Context) error

func NewTransaction(ctx context.Context, fn TransactionFunc) error {
	db := contextPkg.DatabaseFromCtx(ctx)
	return transaction(ctx, db, fn)
}

func transaction(ctx context.Context, db *gorm.DB, fn TransactionFunc) error {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		}
	}()

	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
