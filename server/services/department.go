package services

import (
	"context"
	"log"

	"github.com/awoelf/go-retail/server/config"
	"github.com/awoelf/go-retail/server/graph/model"
)

type Department struct {
	model.Department
}

func (d *Department) AddDepartment(ctx context.Context, input *model.NewDepartment) (int64, error) {
	stmt, err := config.DB.Prepare("INSERT INTO Departments(Name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input.Name)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id, nil
}

func (d *Department) GetAllDepartments(ctx context.Context) ([]*model.Department, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Departments")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var departments []*model.Department

	for res.Next() {
		var department model.Department
		err := res.Scan(&department.ID, &department.Name, &department.TotalSalesDept, &department.CreatedAt, &department.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		departments = append(departments, &department)
	}

	return departments, nil
}

func (d *Department) GetDepartmentById(ctx context.Context, id int64) (*model.Department, error) {
	stmt, err := config.DB.Prepare("SELECT * FROM Departments WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.QueryContext(ctx, id)
	defer res.Close()

	var department model.Department

	for res.Next() {
		err = res.Scan(&department.ID, &department.Name, &department.TotalSalesDept, &department.CreatedAt, &department.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &department, nil
}

func (d *Department) UpdateDepartment(ctx context.Context, input *model.UpdateDepartment) (int64, error) {
	stmt, err := config.DB.Prepare("UPDATE Departments SET Name = ?, TotalSalesDept = ?, UpdatedAt = NOW() WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.ExecContext(ctx, input.Name, input.TotalSalesDept, input.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (d *Department) DeleteDepartment(ctx context.Context, id int64) (error) {
	stmt, err := config.DB.Prepare("DELETE FROM Departments WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Fatal(err)
	}


	return nil
}