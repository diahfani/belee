package products

import "time"

type AddProducts struct {
	BarangName string    `json:"name"`
	Qty        int       `json:"Qty"`
	Price      int       `json:"harga"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
