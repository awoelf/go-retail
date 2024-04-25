package services

import (
	"context"
	"time"

	"github.com/awoelf/go-retail/graph/model"
)

type Department struct {
	model.Department
}

func (d *Department) AddDepartment(ctx context.Context, input *model.NewDepartment) (*model.NewDepartment, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `
		INSERT INTO Departments(
			Name
			CreatedAt
			UpdatedAt
		) 
		VALUES($1, $2, $2)
	`

	_, err := db.ExecContext(ctx, query, input.Name, time.Now())
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (d *Department) GetAllDepartments(ctx context.Context) ([]*model.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `SELECT * FROM Departments ORDER BY ID`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(
			&department.ID, 
			&department.Name, 
			&department.TotalSalesDept, 
			&department.CreatedAt, 
			&department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}

	return departments, nil
}

func (d *Department) GetDepartmentById(ctx context.Context, id *string) (*model.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `SELECT * FROM Departments WHERE ID = $1`

	row, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var department model.Department
	err = row.Scan(
		&department.ID, 
		&department.Name, 
		&department.TotalSalesDept, 
		&department.CreatedAt, 
		&department.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return &department, nil
}

func (d *Department) UpdateDepartment(ctx context.Context, input *model.UpdateDepartment) (*model.UpdateDepartment, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `
		UPDATE Departments 
		SET 
			Name = $1, 
			UpdatedAt = $2 
		WHERE ID = $3
	`

	_, err := db.ExecContext(ctx, query, input.Name, time.Now(), input.ID)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (d *Department) DeleteDepartment(ctx context.Context, id *string) (error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `DELETE FROM Departments WHERE ID = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Department) GetTopDepartments(ctx context.Context) ([]*model.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `SELECT * FROM Departments ORDER BY TotalSalesDept DESC`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(
			&department.ID, 
			&department.Name, 
			&department.TotalSalesDept, 
			&department.CreatedAt, 
			&department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}

	return departments, nil
}