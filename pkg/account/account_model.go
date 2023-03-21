package account

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model `json:"-"`
	ID         uint      `gorm:"primaryKey"  json:"id"`
	Name       string    `gorm:"size:255"  json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
