package graph

import "github.com/awoelf/go-retail/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// TODO Add item, department, and manager services
type Resolver struct{
	Services services.Services
}
