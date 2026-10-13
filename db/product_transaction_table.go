package db

import (
	"time"
	"uuid"
)

type ProductTransactionTable []*ProductTransactionRow

type ProductTransactionRow struct {
	transactionId uuid.UUID

	code OperationCode

	detail string

	createdTime time.Time
}

func NewProductTransactionRow(code OperationCode, detail string) *ProductTransactionRow {
	return &ProductTransactionRow{code: code, detail: detail, transactionId: uuid.NewV4(), createdTime: time.Now()}
}
