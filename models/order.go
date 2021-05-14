package models

type Order struct {
	BaseModel
	CustomerName string `json:"customer_name"`
	Items        []Item `gorm:"foreignKey:OrderId;references:ID" json:"items"`
}
