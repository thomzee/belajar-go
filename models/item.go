package models

import "github.com/google/uuid"

type Item struct {
	BaseModel
	Code    string    `json:"code"`
	Desc    string    `json:"desc"`
	Qty     uint      `json:"qty"`
	OrderId uuid.UUID `json:"order_id"`
}
