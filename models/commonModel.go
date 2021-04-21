package models

import "time"

// common model fields
type CommonModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	// DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
