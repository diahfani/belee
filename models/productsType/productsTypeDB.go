package productsType

import (
	"time"
)

type ProductsType struct {
	Id int `json:"id" gorm:"primaryKey"`
	// ProductsID []products.Products `gorm:"foreignKey:BarangTypeID"`
	NameType  string    `json:"nametype"`
	CreatedAt time.Time `json:"createdAt"`
}
