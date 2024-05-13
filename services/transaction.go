package services

import (
	"context"
	"time"

	"github.com/awoelf/go-retail/graph/model"
)

type Transaction struct{}

func (t *Transaction) SellTransaction(ctx context.Context, input *model.NewTransaction) (*model.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `
		INSERT INTO transactions(
			type,
			paymentMethod,
			items,
			qtyItems,
			totalCost,
			savings,
			status,
			createdAt,
			updatedAt
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $8)
	`

	

	_, err := db.ExecContext(ctx, query, input.Type, input.PaymentMethod, input.Items, time.Now())
	if err != nil {
		return nil, err
	}

	return &model.Transaction{

	}, nil
}

// func (t *Transaction) ReturnTransaction(ctx context.Context, input *model.ItemTransaction) (*model.ItemTransaction, error) {
// 	ctx, cancel := context.WithTimeout(ctx, Timeout)
// 	defer cancel()

// 	query := `UPDATE Items SET
// 		Qty = (Qty + $1),
// 		QtySold = (QtySold - $1),
// 		UpdatedAt = $2
// 		WHERE ID = $3
// 	`

// 	_, err := db.ExecContext(ctx, query, input.QtyTransaction, time.Now(), input.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return input, nil
// }