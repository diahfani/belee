package warung

import (
	"final_project/belee/models/products"
	"time"
)

type Warungs struct {
	Id         int                 `gorm:"primaryKey" json:"id"`
	OwnersID   int                 `json:"ownersId" gorm:"not null"`
	Products   []products.Products `gorm:"foreignKey:WarungID"`
	Name       string              `json:"name" gorm:"not null"`
	Address    string              `json:"address" gorm:"not null"`
	OwnersName string              `json:"ownersName" gorm:"not null"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}