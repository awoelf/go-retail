package services

import (
	"context"
	"time"

	"github.com/awoelf/go-retail/graph/model"
)

type Manager struct{}

func (m *Manager) AddManager(ctx context.Context, input *model.NewManager) (*model.NewManager, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `
		INSERT INTO managers(
			firstName, 
			lastName, 
			departmentID,
			createdAt,
			updatedAt
		) 
		VALUES($1,$2,$3,$4,$4)
	`

	_, err := db.ExecContext(ctx, query, input.FirstName, input.LastName, input.DepartmentID, time.Now())
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (m *Manager) GetAllManagers(ctx context.Context) ([]*model.Manager, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `SELECT * FROM managers`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var managers []*model.Manager
	for rows.Next() {
		var manager model.Manager
		err := rows.Scan(
			&manager.ID,
			&manager.FirstName,
			&manager.LastName,
			&manager.DepartmentID,
			&manager.CreatedAt,
			&manager.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		managers = append(managers, &manager)
	}

	return managers, nil
}

func (m *Manager) GetManagerById(ctx context.Context, id *string) (*model.Manager, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `SELECT * FROM managers WHERE id = $1`

	row, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var manager model.Manager
	err = row.Scan(
		&manager.ID, 
		&manager.FirstName, 
		&manager.LastName, 
		&manager.DepartmentID, 
		&manager.CreatedAt, 
		&manager.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return &manager, nil
}

func (m *Manager) UpdateManager(ctx context.Context, input *model.UpdateManager) (*model.UpdateManager, error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `
		UPDATE managers 
		SET 
			firstName = $1, 
			lastName = $2, 
			departmentID = $3, 
			updatedAt = $4 
		WHERE id = $5
	`

	_, err := db.ExecContext(
		ctx,
		query,
		input.FirstName, 
		input.LastName, 
		input.DepartmentID, 
		time.Now(),
		input.ID,
	)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (m *Manager) DeleteManager(ctx context.Context, id *string) (error) {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()
	
	query := `DELETE FROM managers WHERE id = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}