package entity

// Product : Attribute at product table
type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
}
