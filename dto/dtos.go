package dto

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type Order struct {
	ID      string  `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}
