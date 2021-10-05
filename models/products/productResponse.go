package products

import (
	"belee/models/productsType"
	"belee/models/warung"
	"time"
)

type ProductResponse struct {
	Id         int                        `json:"id"`
	Warung     *warung.Warungs            `json:"warung"`
	BarangType *productsType.ProductsType `json:"barangType"`
	BarangName string                     `json:"name"`
	Qty        int                        `json:"qty"`
	Price      int                        `json:"price"`
	CreatedAt  time.Time                  `json:"created_at"`
	UpdatedAt  time.Time                  `json:"updated_at"`
}
