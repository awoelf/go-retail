package services

import (
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Manager struct {
	model.Manager
}

func (m *Manager) AddManager(input *model.NewManager) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Managers(FirstName, LastName, DepartmentID) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.FirstName, input.LastName, input.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (m *Manager) GetAllManagers() ([]*model.Manager, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Managers")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var managers []*model.Manager

	for res.Next() {
		var manager model.Manager
		err := res.Scan(&manager.ID, &manager.FirstName, &manager.LastName, &manager.DepartmentID, &manager.CreatedAt, &manager.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		managers = append(managers, &manager)
	}

	return managers, nil
}

func (m *Manager) GetManagerById(id int64) (*model.Manager, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Managers WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query(id)
	defer res.Close()

	var manager model.Manager

	for res.Next() {
		err = res.Scan(&manager.ID, &manager.FirstName, &manager.LastName, &manager.DepartmentID, &manager.CreatedAt, &manager.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &manager, nil
}

func (m *Manager) UpdateManager(input *model.UpdateManager) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Managers SET FirstName = ?, LastName = ?, DepartmentID = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.FirstName, input.LastName, input.DepartmentID, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (m *Manager) DeleteManager(id int64) (error) {
	stmt, err := config.DB.Prepare("DELETE FROM Managers WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}


	return nil
}