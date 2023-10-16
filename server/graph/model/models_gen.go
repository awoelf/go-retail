// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Aisle struct {
	ID           int     `json:"id"`
	DepartmentID string  `json:"departmentId"`
	CreatedAt    *string `json:"createdAt,omitempty"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
}

type Department struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	TotalSalesWeekDept *float64 `json:"totalSalesWeekDept,omitempty"`
	CreatedAt          *string  `json:"createdAt,omitempty"`
	UpdatedAt          *string  `json:"updatedAt,omitempty"`
}

type Item struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Price              float64  `json:"price"`
	Qty                int      `json:"qty"`
	Category           string   `json:"category"`
	Promotion          *bool    `json:"promotion,omitempty"`
	PromotionPrice     *float64 `json:"promotionPrice,omitempty"`
	Replenish          *bool    `json:"replenish,omitempty"`
	TotalSalesWeekItem *float64 `json:"totalSalesWeekItem,omitempty"`
	AisleID            string   `json:"aisleId"`
	DepartmentID       string   `json:"departmentId"`
	CreatedAt          *string  `json:"createdAt,omitempty"`
	UpdatedAt          *string  `json:"updatedAt,omitempty"`
}

type Manager struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	DepartmentID *string `json:"departmentId,omitempty"`
	CreatedAt    *string `json:"createdAt,omitempty"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
}

type NewAisle struct {
	DepartmentID string `json:"departmentId"`
}

type NewDepartment struct {
	Name string `json:"name"`
}

type NewItem struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Qty          int     `json:"qty"`
	Category     string  `json:"category"`
	AisleID      string  `json:"aisleId"`
	DepartmentID string  `json:"departmentId"`
}

type NewManager struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DepartmentID string `json:"departmentId"`
}

type UpdateAisle struct {
	ID           int    `json:"id"`
	DepartmentID string `json:"departmentId"`
}

type UpdateDepartment struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	TotalSalesWeekDept float64 `json:"totalSalesWeekDept"`
}

type UpdateItem struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	Qty                int     `json:"qty"`
	Category           string  `json:"category"`
	Promotion          bool    `json:"promotion"`
	PromotionPrice     float64 `json:"promotionPrice"`
	Replenish          bool    `json:"replenish"`
	TotalSalesWeekItem float64 `json:"totalSalesWeekItem"`
	AisleID            string  `json:"aisleId"`
	DepartmentID       string  `json:"departmentId"`
}

type UpdateManager struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DepartmentID string `json:"departmentId"`
}
