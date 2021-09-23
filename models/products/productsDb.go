package products

import (
	"time"
)

type Products struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	WarungID     uint      `json:"warungId"`
	BarangTypeID int       `json:"barangTypeId"`
	WarungName   int       `json:"warungName"`
	ProductsName string    `json:"productsName"`
	Qty          string    `json:"qty"`
	Price        int       `json:"price"`
	Created_at   time.Time `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP"`
	Updated_at   time.Time `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP"`
}
