package models

type Order struct {
	OrderId    uint   `gorm:"primaryKey" json:"id,omitempty"`
	OrderName  string `json:"order_name"`
	UserId     int    `json:"user_id"`
	User       User   `json:"user"`
	Invoice    string `json:"invoice"`
	Address    string `json:"address"`
	Telphone   string `json:"telphone"`
	Amount     int    `json:"amount"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}

type Orders []Order
