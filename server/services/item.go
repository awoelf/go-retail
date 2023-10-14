package services

import (
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Item struct{ model.Item }

func (item Item) AddItem() int64 {
	stmt, err := config.DB.Prepare("INSERT INTO Items(Name, Price, Qty, Category, AisleID, DepartmentID) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(item.Name, item.Price, item.Qty, item.Category, item.AisleID, item.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Item created!")
	return id
}
