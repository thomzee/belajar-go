package dto

type OrderDTO struct {
	CustomerName string `json:"customer_name"`
	Items        []ItemDTO `json:"items"`
}

type ItemDTO struct {
	Code    string `json:"code"`
	Desc    string `json:"desc"`
	Qty     uint   `json:"qty"`
}