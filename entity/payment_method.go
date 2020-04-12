package entity

// PaymentMethod : Attribute at payment_method table
type PaymentMethod struct {
	ID            int    `json:"id"`
	PaymentName   string `json:"payment_name"`
	AccountNumber int    `json:"account_number"`
}
