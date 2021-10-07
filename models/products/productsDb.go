package products

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	WarungID     int       `json:"warungId"`
	BarangTypeID int       `json:"barangtypeId"`
	BarangName   string    `json:"productsName" gorm:"not null"`
	Qty          int       `json:"qty"`
	Price        int       `json:"price"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt    gorm.DeletedAt
}
