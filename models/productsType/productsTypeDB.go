package productsType

import "time"

type ProductsType struct {
	Id         int `json:"id" gorm:"primaryKey"`
	ProductsID string
	NameType   string    `json:"nametype"`
	CreatedAt  time.Time `json:"createdAt"`
}
