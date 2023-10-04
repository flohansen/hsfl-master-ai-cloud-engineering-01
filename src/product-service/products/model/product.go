package model

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Retailer    string  `json:"retailer"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
