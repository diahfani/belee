package warung

import (
	"time"

	"gorm.io/gorm"
)

type Warungs struct {
	Id int `gorm:"primaryKey" json:"id"`
	// OwnersID int                 `json:"ownersId"`
	// Products []products.Products `gorm:"foreignKey:WarungID"`
	Name    string `json:"name"`
	Address string `json:"address"`
	// OwnersName string              `json:"ownersName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   gorm.DeletedAt
}
