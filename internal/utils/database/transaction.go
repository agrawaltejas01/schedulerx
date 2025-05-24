package database_utils

import (
	"context"
	"fmt"

	db "github.com/agrawaltejas01/schedulerx/internal/database"
	"gorm.io/gorm"
)

type transactionKeyType string

const (
	TransactionKey transactionKeyType = "db_transaction"
)

func setTxnInContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TransactionKey, tx)
}

func getTxnFromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(TransactionKey).(*gorm.DB)
	if !ok {
		fmt.Println("No transaction found in context")
		return nil
	}
	return tx
}

func GetTxnOrDb(ctx context.Context) *gorm.DB {
	// Check if a transaction exists in the context
	if tx := getTxnFromContext(ctx); tx != nil {
		return tx
	}

	// If no transaction exists, return the default database connection
	return db.Database
}

func StartTransaction(ctx context.Context) context.Context {

	// Check if a transaction already exists in the context
	if tx := getTxnFromContext(ctx); tx != nil {
		fmt.Println("Transaction already exists in context")
		return ctx
	}

	tx := db.Database.Begin()
	if tx.Error != nil {
		fmt.Println("Error starting transaction in Repo Layer")
		return nil
	}
	return setTxnInContext(ctx, tx)
}

func EndTransaction(ctx context.Context, err error) error {

	tx := getTxnFromContext(ctx)
	if tx == nil {
		fmt.Println("No transaction found in context to end")
		return fmt.Errorf("no transaction found in context")
	}

	if err == nil {
		err := tx.Commit().Error
		if err != nil {
			fmt.Println("Error in Committing TXN")
			return err
		}
	} else {

		err = tx.Rollback().Error
		if err != nil {
			fmt.Println("Error in Rolling Back TXN")
			return err
		}
	}
	return nil
}
