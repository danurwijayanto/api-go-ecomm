package entity

// Transaction : Attribute at transaction table
type Transaction struct {
	ID            int    `json:"id"`
	IDProduct     int    `json:"product_id"`
	IDCategory    int    `json:"category_id"`
	CodeNumber    string `json:"code_number"`
	PhoneNumber   string `json:"phone_number"`
	InvoiceCode   string `json:"invoice_code"`
	Nominal       string `json:"nominal"`
	Status        int    `json:"status"`
	PaymentMethod int    `json:"payment_method"`
}
