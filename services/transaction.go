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

	queryTransaction := `
		INSERT INTO transactions(
			type,
			paymentMethod,
			items,
			qtyItems,
			totalCost,
			savings,
			createdAt,
			updatedAt
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $7)
	`
	
	queryItems := `SELECT * FROM items WHERE id = ANY($1)`
	rows, queryErr := db.QueryContext(ctx, queryItems)
	if queryErr != nil {
		return nil, queryErr
	}

	var qtyItems = len(input.Items)
	var totalCost float64
	var savings float64
	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.ID,
			&item.DepartmentID,
			&item.Name,
			&item.Price,
			&item.Qty,
			&item.QtySold,
			&item.Category,
			&item.Promo,
			&item.PromoPrice,
			&item.TotalSalesItem,
			&item.Aisle,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		if (*item.Promo) {
			totalCost += *item.PromoPrice
			savings += *item.PromoPrice
		} else {
			totalCost += item.Price
		}
	}
	
	_, err := db.ExecContext(
		ctx, 
		queryTransaction, 
		input.Type, 
		input.PaymentMethod, 
		input.Items, 
		qtyItems, 
		totalCost, 
		savings, 
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return &model.Transaction{
		Type: input.Type,
		PaymentMethod: input.PaymentMethod,
		Items: input.Items,
		QtyItems: qtyItems,
		TotalCost: totalCost,
		Savings: savings,
	}, nil
}

func (t *Transaction) ReturnTransaction(ctx context.Context, input *model.NewTransaction) (*model.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	queryTransaction := `
		UPDATE transactions 
		SET
			type = $1,
			items = $2,
			qtyItems = (qtyItems - $3),
			totalCost = (totalCost - $4),
			savings = (savings - $5),
			updatedAt = $6
		WHERE id = $7
	`
	
	queryItems := `SELECT * FROM items WHERE id = ANY($1)`
	rows, queryErr := db.QueryContext(ctx, queryItems)
	if queryErr != nil {
		return nil, queryErr
	}

	var qtyItemsReturned = len(input.Items)
	var totalCostReturned float64
	var savingsReturned float64
	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.ID,
			&item.DepartmentID,
			&item.Name,
			&item.Price,
			&item.Qty,
			&item.QtySold,
			&item.Category,
			&item.Promo,
			&item.PromoPrice,
			&item.TotalSalesItem,
			&item.Aisle,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		if (*item.Promo) {
			totalCostReturned += *item.PromoPrice
			savingsReturned += *item.PromoPrice
		} else {
			totalCostReturned += item.Price
		}
	}
	
	_, err := db.ExecContext(
		ctx, 
		queryTransaction, 
		input.Type, 
		input.Items, 
		qtyItemsReturned, 
		totalCostReturned, 
		savingsReturned, 
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return &model.Transaction{
		Type: input.Type,
		PaymentMethod: input.PaymentMethod,
		Items: input.Items,
	}, nil
}