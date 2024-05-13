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
		INSERT INTO items(
			departmentID, 
			name, 
			price, 
			qty, 
			category, 
			aisle, 
			createdAt, 
			updatedAt
		) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$7)
	`

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
	)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM items ORDER BY id`

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

func (i *Item) GetItemById(ctx context.Context, id *string) (*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM items WHERE id = $1`

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

	query := `
		UPDATE items 
		SET 
			name = $1,
			departmentID = $2,
			price = $3, 
			qty = $4, 
			qtySold = $5
			category = $6, 
			promo = $7,
			promoPrice = $8
			totalSalesItem = $9, 
			aisle = $10,
			updatedAt = $11 
		WHERE id = $12
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

	query := `DELETE FROM items WHERE id = ?`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (i *Item) OrderItems(ctx context.Context, input *model.ItemOrder) (*model.ItemOrder, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE items SET qty = (qty + $1), updatedAt = $2 WHERE id = $3`

	_, err := db.ExecContext(ctx, query, input.QtyOrder, time.Now(), input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) GetTopItems(ctx context.Context) ([]*model.Item, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `SELECT * FROM items ORDER BY qtySold DESC`

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

	query := `SELECT * FROM items WHERE category = $1`

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

	query := `UPDATE items SET promo = $1, promoPrice = $2 WHERE id = $3`

	_, err := db.ExecContext(ctx, query, input.Promo, input.PromoPrice, input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *Item) EndSaleItem(ctx context.Context, id *string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	query := `UPDATE items SET promo = false WHERE id = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}
