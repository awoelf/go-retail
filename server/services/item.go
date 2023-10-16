package services

import (
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Item struct{ model.Item }

func (i *Item) AddItem(input *model.NewItem) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Items(Name, Price, Qty, Category, AisleID) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.Name, input.Price, input.Qty, input.Category, input.AisleID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (i *Item) GetAllItems() ([]*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var items []*model.Item

	for res.Next() {
		var item model.Item
		err := res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.PromotionPrice, &item.Replenish, &item.TotalSalesWeekItem, &item.AisleID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	return items, nil
}

func (i *Item) GetItemById(id int64) (*model.Item, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Items WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query(id)
	defer res.Close()

	var item model.Item

	for res.Next() {
		err = res.Scan(&item.ID, &item.Name, &item.Price, &item.Qty, &item.Category, &item.Promotion, &item.PromotionPrice, &item.Replenish, &item.TotalSalesWeekItem, &item.AisleID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &item, nil
}

func (i *Item) UpdateItem(input *model.UpdateItem) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Items SET Name = ?, Price = ?, Qty = ?, Category = ?, Promotion = ?, PromotionPrice = ?, Replenish = ?, TotalSalesWeekItem = ?, AisleID = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.Name, input.Price, input.Qty, input.Category, input.Promotion, input.PromotionPrice, input.Replenish, input.TotalSalesWeekItem, input.AisleID, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (i *Item) DeleteItem(id int64) (error) {
	stmt, err := config.DB.Prepare("DELETE FROM Items WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}


	return nil
}