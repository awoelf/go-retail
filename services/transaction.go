package services

import (
	"context"

	"github.com/awoelf/go-retail/graph/model"
)

type Transaction struct{}

func (t *Transaction) SellTransaction(ctx context.Context, input *model.ItemTransaction) (*model.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `
		INSERT INTO transactions(
		
		)

	`

	_, err := db.ExecContext(ctx, query, input.QtyTransaction, time.Now(), input.ID)
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

// Equates to scanning an item during transaction in person or adding item to cart online
func (t *Transaction) AddItemTransaction() {

}

func (t *Transaction) RemoveItemTransaction() {

}