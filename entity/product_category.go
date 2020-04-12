package entity

// ProductCategory : Attribute at product_category table
type ProductCategory struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	ProductID    int    `json:"product_id"`
}
