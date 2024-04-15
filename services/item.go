package services

import (
	"context"
	"log"

	"github.com/awoelf/go-retail/config"
	"github.com/awoelf/go-retail/graph/model"
)

type Item struct{ model.Item }

func (i *Item) AddItem(ctx context.Context, input *model.NewItem) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Items(Name, Price, Qty, Category, Aisle, DepartmentID) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input.Name, input.Price, input.Qty, input.Category, input.Aisle, input.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (i *Item) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var items []*model.Item

	for res.Next() {
		var item model.Item
		err := res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.TotalSalesItem, &item.Aisle, &item.DepartmentID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) GetItemById(ctx context.Context, id int64) (*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx, id)
	defer res.Close()

	var item model.Item

	for res.Next() {
		err = res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.TotalSalesItem, &item.Aisle, &item.DepartmentID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &item, nil
}

func (i *Item) UpdateItem(ctx context.Context, input *model.UpdateItem) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Items SET Name = ?, Price = ?, Qty = ?, Category = ?, Promotion = ?, TotalSalesItem = ?, Aisle = ?, DepartmentID = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input.Name, input.Price, input.Qty, input.Category, input.Promotion, input.TotalSalesItem, input.Aisle, input.DepartmentID, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) DeleteItem(ctx context.Context, id int64) error {
	stmt, err := config.DB.Prepare("DELETE FROM Items WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (i *Item) SellItem(ctx context.Context, input *model.ItemTransaction) (int64, error) {
	var totalSales = input.Price * float64(input.QtySold)

	// Add sales to department total
	stmtDept, err := config.DB.Prepare("UPDATE Departments SET TotalSalesDept = (TotalSalesDept + ?) WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmtDept.ExecContext(ctx, totalSales, input.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	// Add sales to item total and remove quantity from item
	stmtItem, err := config.DB.Prepare("UPDATE Items SET Qty = (Qty - ?), QtySold = (QtySold + ?), TotalSalesItem = (TotalSalesItem + ?) WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmtItem.ExecContext(ctx, input.QtySold, input.QtySold, totalSales, input.ID)
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) ReturnItem(ctx context.Context, input *model.ItemTransaction) (int64, error) {
	var totalSales = input.Price * float64(input.QtySold)

	// Remove sales from department total
	stmtDept, err := config.DB.Prepare("UPDATE Departments SET TotalSalesDept = (TotalSalesDept - ?) WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmtDept.ExecContext(ctx, totalSales, input.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	// Remove sales from item total and add quantity to item
	stmtItem, err := config.DB.Prepare("UPDATE Items SET Qty = (Qty + ?), QtySold = (QtySold - ?), TotalSalesItem = (TotalSalesItem - ?) WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmtItem.ExecContext(ctx, input.QtySold, input.QtySold, totalSales, input.ID)
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) OrderItems(ctx context.Context, input *model.ItemOrder) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Items SET Qty = (Qty + ?) WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input.Qty, input.ID)
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) GetTopItems(ctx context.Context) ([]*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items ORDER BY QtySold DESC")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var items []*model.Item

	for res.Next() {
		var item model.Item
		err := res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.TotalSalesItem, &item.Aisle, &item.DepartmentID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) GetItemsByCategory(ctx context.Context, category *string) ([]*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items WHERE Category = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx, category)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var items []*model.Item

	for res.Next() {
		var item model.Item
		err := res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.TotalSalesItem, &item.Aisle, &item.DepartmentID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) SetSaleItem(ctx context.Context, input *model.ItemPromotion) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Items SET Price = ?, Promotion = true WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) ResetItem(ctx context.Context, input *model.ItemPromotion) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Items SET Price = ?, Promotion = false WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}