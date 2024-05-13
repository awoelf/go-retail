// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Department struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	TotalSalesDept *float64 `json:"totalSalesDept,omitempty"`
	CreatedAt      *string  `json:"createdAt,omitempty"`
	UpdatedAt      *string  `json:"updatedAt,omitempty"`
}

type Item struct {
	ID             string   `json:"id"`
	DepartmentID   string   `json:"departmentId"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	Qty            int      `json:"qty"`
	QtySold        *int     `json:"qtySold,omitempty"`
	Category       *string  `json:"category,omitempty"`
	Promo          *bool    `json:"promo,omitempty"`
	PromoPrice     *float64 `json:"promoPrice,omitempty"`
	TotalSalesItem *float64 `json:"totalSalesItem,omitempty"`
	Aisle          string   `json:"aisle"`
	CreatedAt      *string  `json:"createdAt,omitempty"`
	UpdatedAt      *string  `json:"updatedAt,omitempty"`
}

type ItemOrder struct {
	ID       string `json:"id"`
	QtyOrder int    `json:"qtyOrder"`
}

type ItemPromotion struct {
	ID         string  `json:"id"`
	Promo      bool    `json:"promo"`
	PromoPrice float64 `json:"promoPrice"`
}

type Manager struct {
	ID           string  `json:"id"`
	DepartmentID *string `json:"departmentId,omitempty"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	CreatedAt    *string `json:"createdAt,omitempty"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
}

type Mutation struct {
}

type NewDepartment struct {
	Name string `json:"name"`
}

type NewItem struct {
	DepartmentID   string   `json:"departmentId"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	Qty            int      `json:"qty"`
	QtySold        *int     `json:"qtySold,omitempty"`
	Category       *string  `json:"category,omitempty"`
	Promo          *bool    `json:"promo,omitempty"`
	PromoPrice     *float64 `json:"promoPrice,omitempty"`
	TotalSalesItem *float64 `json:"totalSalesItem,omitempty"`
	Aisle          string   `json:"aisle"`
}

type NewManager struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DepartmentID string `json:"departmentId"`
}

type NewTransaction struct {
	Type          string    `json:"type"`
	PaymentMethod string    `json:"paymentMethod"`
	Items         []*string `json:"items"`
}

type Query struct {
}

type Transaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	PaymentMethod string    `json:"paymentMethod"`
	Items         []*string `json:"items"`
	QtyItems      int       `json:"qtyItems"`
	TotalCost     float64   `json:"totalCost"`
	Savings       float64   `json:"savings"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
}

type UpdateDepartment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateItem struct {
	ID             string   `json:"id"`
	DepartmentID   *string  `json:"departmentId,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Price          *float64 `json:"price,omitempty"`
	Qty            *int     `json:"qty,omitempty"`
	QtySold        *int     `json:"qtySold,omitempty"`
	Category       *string  `json:"category,omitempty"`
	Promo          *bool    `json:"promo,omitempty"`
	PromoPrice     *float64 `json:"promoPrice,omitempty"`
	TotalSalesItem *float64 `json:"totalSalesItem,omitempty"`
	Aisle          *string  `json:"aisle,omitempty"`
}

type UpdateManager struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DepartmentID string `json:"departmentId"`
}
