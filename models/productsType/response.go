package productsType

import "time"

type ProductsTypeResponse struct {
	Id        int       `json:"id"`
	NameType  string    `json:"nametype"`
	CreatedAt time.Time `json:"createdAt"`
}
