package services

import (
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Aisle struct {
	model.Aisle
}

func (a *Aisle) AddAisle(input *model.NewAisle) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Aisles(DepartmentID) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (a *Aisle) GetAllAisles() ([]*model.Aisle, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Aisles")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var aisles []*model.Aisle

	for res.Next() {
		var aisle model.Aisle
		err := res.Scan(&aisle.ID, &aisle.DepartmentID, &aisle.CreatedAt, &aisle.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		aisles = append(aisles, &aisle)
	}

	return aisles, nil
}

func (a *Aisle) GetAisleById(id int64) (*model.Aisle, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Aisles WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query(id)
	defer res.Close()

	var aisle model.Aisle

	for res.Next() {
		err = res.Scan(&aisle.ID, &aisle.DepartmentID, &aisle.CreatedAt, &aisle.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &aisle, nil
}

func (a *Aisle) UpdateAisle(input *model.UpdateAisle) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Aisles SET DepartmentID = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.DepartmentID, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (a *Aisle) DeleteAisle(id int64) (error) {
	stmt, err := config.DB.Prepare("DELETE FROM Aisles WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}


	return nil
}