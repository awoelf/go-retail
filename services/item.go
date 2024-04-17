package services

import (
	"context"
	"log"
	"time"

	"github.com/awoelf/go-retail/graph/model"
)

type Item struct{}

func (i *Item) AddItem(ctx context.Context, input *model.NewItem) (*model.NewItem, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `
		INSERT INTO Items(
			DepartmentID, 
			Name, 
			Price, 
			Qty, 
			Category, 
			Aisle, 
			CreatedAt, 
			UpdatedAt
		) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) returning *`

	_, err := db.ExecContext(
		ctx,
		query,
		input.DepartmentID,
		input.Name,
		input.Price,
		input.Qty,
		input.Category,
		input.Aisle,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM Items ORDER BY ID`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var items []*model.Item
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
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) GetItemById(ctx context.Context, id string) (*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM Items WHERE ID = $1`

	row := db.QueryRowContext(ctx, query, id)

	var item model.Item
	err := row.Scan(
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

	return &item, nil
}

func (i *Item) UpdateItem(ctx context.Context, input *model.UpdateItem) (*model.UpdateItem, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items 
		SET 
			Name = $1,
			DepartmentID = $2,
			Price = $3, 
			Qty = $4, 
			QtySold = $5
			Category = $6, 
			Promo = $7,
			PromoPrice = $8
			TotalSalesItem = $9, 
			Aisle = $10,
			UpdatedAt = $11 
		WHERE ID = $12
		returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		input.Name,
		input.DepartmentID,
		input.Price,
		input.Qty,
		input.QtySold,
		input.Category,
		input.Promo,
		input.PromoPrice,
		input.TotalSalesItem,
		input.Aisle,
		time.Now(),
		input.ID,
	)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) DeleteItem(ctx context.Context, id *string) error {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `DELETE FROM Items WHERE ID = ?`

	_, err := db.ExecContext(ctx, query, &id)
	if err != nil {
		return err
	}

	return nil
}

func (i *Item) SellItem(ctx context.Context, input *model.ItemTransaction) (*model.ItemTransaction, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items SET 
		Qty = (Qty - $1), 
		QtySold = (QtySold + $1),
		UpdatedAt = $2
		WHERE ID = $3
	`

	_, err := db.ExecContext(ctx, query, input.QtyTransaction, time.Now(), input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) ReturnItem(ctx context.Context, input *model.ItemTransaction) (*model.ItemTransaction, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items SET 
		Qty = (Qty + $1), 
		QtySold = (QtySold - $1),
		UpdatedAt = $2
		WHERE ID = $3
	`

	_, err := db.ExecContext(ctx, query, input.QtyTransaction, time.Now(), input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) OrderItems(ctx context.Context, input *model.ItemOrder) (*model.ItemOrder, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items SET Qty = (Qty + $1), UpdatedAt = $2 WHERE ID = $3`

	_, err := db.ExecContext(ctx, query, input.QtyOrder, time.Now(), input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) GetTopItems(ctx context.Context) ([]*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM Items ORDER BY QtySold DESC`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var items []*model.Item
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
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) GetItemsByCategory(ctx context.Context, category *string) ([]*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM Items WHERE Category = $1`

	rows, err := db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}

	var items []*model.Item
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
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) StartSaleItem(ctx context.Context, input *model.ItemPromotion) (*model.ItemPromotion, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items SET Promo = $1, PromoPrice = $2 WHERE ID = $3`

	_, err := db.ExecContext(ctx, query, input.Promo, input.PromoPrice, input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) EndSaleItem(ctx context.Context, id *string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE Items SET Promotion = false WHERE ID = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}
