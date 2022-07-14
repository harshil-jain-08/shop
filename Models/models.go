package Models

/*type Customer struct {
	ID   int    `json:"customer_id"`
	Name string `json:"customer_name"`
}*/

type Product struct {
	ID       int    `json:"product_id"`
	Name     string `json:"product_name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type Order struct {
	ID         int `json:"order_id"`
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

/*type SubOrder struct {
	gorm.Model
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}*/
