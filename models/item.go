package models

type Item struct {
	CommonModel
	ItemCode    string `gorm:"not null" json:"itemCode"`
	Description string `gorm:"not null" json:"description"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	OrderID     uint   `gorm:"not null" json:"orderId"`
}
