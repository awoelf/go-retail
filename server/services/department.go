package services

import (
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Department struct {
	model.Department
}

func (d *Department) AddDepartment(input *model.NewDepartment) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Departments(Name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.Name)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (d *Department) GetAllDepartments() ([]*model.Department, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Departments")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var departments []*model.Department

	for res.Next() {
		var department model.Department
		err := res.Scan(&department.ID, &department.Name, &department.TotalSalesWeekDept, &department.CreatedAt, &department.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		departments = append(departments, &department)
	}

	return departments, nil
}

func (d *Department) GetDepartmentById(id int64) (*model.Department, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Departments WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Query(id)
	defer res.Close()

	var department model.Department

	for res.Next() {
		err = res.Scan(&department.ID, &department.Name, &department.TotalSalesWeekDept, &department.CreatedAt, &department.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &department, nil
}

func (d *Department) UpdateDepartment(input *model.UpdateDepartment) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Departments SET Name = ?, TotalSalesWeekDept = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(input.Name, input.TotalSalesWeekDept, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (d *Department) DeleteDepartment(id int64) (error) {
	stmt, err := config.DB.Prepare("DELETE FROM Departments WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}


	return nil
}