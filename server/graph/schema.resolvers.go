package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"
	"log"

	"github.com/awoelf/go-retail/server/graph/model"
	"github.com/awoelf/go-retail/server/services"
)

// AddItem is the resolver for the addItem field.
func (r *mutationResolver) AddItem(ctx context.Context, input *model.NewItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: AddItem - addItem"))
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input *model.UpdateItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: UpdateItem - updateItem"))
}

// DeleteItem is the resolver for the deleteItem field.
func (r *mutationResolver) DeleteItem(ctx context.Context, id *int) (*int, error) {
	panic(fmt.Errorf("not implemented: DeleteItem - deleteItem"))
}

// SellItem is the resolver for the sellItem field.
func (r *mutationResolver) SellItem(ctx context.Context) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: SellItem - sellItem"))
}

// ReturnItem is the resolver for the returnItem field.
func (r *mutationResolver) ReturnItem(ctx context.Context) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: ReturnItem - returnItem"))
}

// OrderItems is the resolver for the orderItems field.
func (r *mutationResolver) OrderItems(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: OrderItems - orderItems"))
}

// SetSaleItem is the resolver for the setSaleItem field.
func (r *mutationResolver) SetSaleItem(ctx context.Context) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: SetSaleItem - setSaleItem"))
}

// AddDepartment is the resolver for the addDepartment field.
func (r *mutationResolver) AddDepartment(ctx context.Context, input *model.NewDepartment) (*model.Department, error) {
	var Department services.Department
	id, err := Department.AddDepartment(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %v\n", id)
	return &model.Department{ID: int(id), Name: input.Name}, nil
}

// UpdateDepartment is the resolver for the updateDepartment field.
func (r *mutationResolver) UpdateDepartment(ctx context.Context, input *model.UpdateDepartment) (*model.Department, error) {
	var Department services.Department
	id, err := Department.UpdateDepartment(input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Department{ID: int(id), Name: input.Name, TotalSalesWeekDept: &input.TotalSalesWeekDept}, nil
}

// DeleteDepartment is the resolver for the deleteDepartment field.
func (r *mutationResolver) DeleteDepartment(ctx context.Context, id *int) (*int, error) {
	var Department services.Department
	err := Department.DeleteDepartment(int64(*id))
	if err != nil {
		log.Fatal(err)
	}
	
	return id, nil
}

// AddAisle is the resolver for the addAisle field.
func (r *mutationResolver) AddAisle(ctx context.Context, input *model.NewAisle) (*model.Aisle, error) {
	var Aisle services.Aisle
	id, err := Aisle.AddAisle(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %v\n", id)
	return &model.Aisle{ID: int(id), DepartmentID: input.DepartmentID}, nil
}

// UpdateAisle is the resolver for the updateAisle field.
func (r *mutationResolver) UpdateAisle(ctx context.Context, input *model.UpdateAisle) (*model.Aisle, error) {
	var Aisle services.Aisle
	id, err := Aisle.UpdateAisle(input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Aisle{ID: int(id), DepartmentID: input.DepartmentID}, nil
}

// DeleteAisle is the resolver for the deleteAisle field.
func (r *mutationResolver) DeleteAisle(ctx context.Context, id *int) (*int, error) {
	var Aisle services.Aisle
	err := Aisle.DeleteAisle(int64(*id))
	if err != nil {
		log.Fatal(err)
	}
	
	return id, nil}

// AddManager is the resolver for the addManager field.
func (r *mutationResolver) AddManager(ctx context.Context, input *model.NewManager) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented: AddManager - addManager"))
}

// UpdateManager is the resolver for the updateManager field.
func (r *mutationResolver) UpdateManager(ctx context.Context, input *model.UpdateManager) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented: UpdateManager - updateManager"))
}

// DeleteManager is the resolver for the deleteManager field.
func (r *mutationResolver) DeleteManager(ctx context.Context, id *int) (*int, error) {
	panic(fmt.Errorf("not implemented: DeleteManager - deleteManager"))
}

// GetAllItems is the resolver for the getAllItems field.
func (r *queryResolver) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: GetAllItems - getAllItems"))
}

// GetItemByID is the resolver for the getItemById field.
func (r *queryResolver) GetItemByID(ctx context.Context, id *int) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: GetItemByID - getItemById"))
}

// GetTopItems is the resolver for the getTopItems field.
func (r *queryResolver) GetTopItems(ctx context.Context) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: GetTopItems - getTopItems"))
}

// GetItemsByCategory is the resolver for the getItemsByCategory field.
func (r *queryResolver) GetItemsByCategory(ctx context.Context, category *string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: GetItemsByCategory - getItemsByCategory"))
}

// GetAllDepartments is the resolver for the getAllDepartments field.
func (r *queryResolver) GetAllDepartments(ctx context.Context) ([]*model.Department, error) {
	var Department services.Department
	var resDepartments []*model.Department
	dbDepartments, err := Department.GetAllDepartments()
	if err != nil {
		log.Fatal(err)
	}

	for _, department := range dbDepartments {
		resDepartments = append(resDepartments, &model.Department{ID: department.ID, Name: department.Name, TotalSalesWeekDept: department.TotalSalesWeekDept, CreatedAt: department.CreatedAt, UpdatedAt: department.UpdatedAt})
	}

	return resDepartments, nil
}

// GetDepartmentByID is the resolver for the getDepartmentById field.
func (r *queryResolver) GetDepartmentByID(ctx context.Context, id *int) (*model.Department, error) {
	var Department services.Department
	department, err := Department.GetDepartmentById(int64(*id))
	if err != nil {
		log.Fatal(err)
	}

	return department, err
}

// GetTopDepartments is the resolver for the getTopDepartments field.
func (r *queryResolver) GetTopDepartments(ctx context.Context) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: GetTopDepartments - getTopDepartments"))
}

// GetAllAisles is the resolver for the getAllAisles field.
func (r *queryResolver) GetAllAisles(ctx context.Context) ([]*model.Aisle, error) {
	var Aisle services.Aisle
	var resAisles []*model.Aisle
	dbAisles, err := Aisle.GetAllAisles()
	if err != nil {
		log.Fatal(err)
	}

	for _, aisle := range dbAisles {
		resAisles = append(resAisles, &model.Aisle{ID: aisle.ID, DepartmentID: aisle.DepartmentID, CreatedAt: aisle.CreatedAt, UpdatedAt: aisle.UpdatedAt})
	}

	return resAisles, nil}

// GetAisleByID is the resolver for the getAisleById field.
func (r *queryResolver) GetAisleByID(ctx context.Context, id *int) (*model.Aisle, error) {
	var Aisle services.Aisle
	aisle, err := Aisle.GetAisleById(int64(*id))
	if err != nil {
		log.Fatal(err)
	}

	return aisle, err}

// GetAllManagers is the resolver for the getAllManagers field.
func (r *queryResolver) GetAllManagers(ctx context.Context) ([]*model.Manager, error) {
	panic(fmt.Errorf("not implemented: GetAllManagers - getAllManagers"))
}

// GetManagerByID is the resolver for the getManagerById field.
func (r *queryResolver) GetManagerByID(ctx context.Context, id *int) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented: GetManagerByID - getManagerById"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
