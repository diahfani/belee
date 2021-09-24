package products

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	WarungID     uint           `json:"warungId"`
	BarangTypeID int            `json:"barangTypeId"`
	WarungName   int            `json:"warungName"`
	BarangName   string         `json:"productsName"`
	Qty          int            `json:"qty"`
	Price        int            `json:"price"`
	CreatedAt    time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
