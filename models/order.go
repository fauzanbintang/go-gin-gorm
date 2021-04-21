package models

import "time"

type Order struct {
	CommonModel
	CustomerName string    `gorm:"not null" json:"customerName"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"foreignKey:OrderID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
