package dto

type Product struct {
	Name     string `json:"product_name"`
	Quantity int    `json:"stock"`
	Price    int    `json:"price"`
}

type Order struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
