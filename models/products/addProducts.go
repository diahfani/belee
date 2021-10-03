package products

import "time"

type AddProducts struct {
	BarangName string    `json:"name"`
	Qty        int       `json:"qty"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
