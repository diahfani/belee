package productsType

type ProductsType struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	NameType string `json:"nametype"`
}

type AddType struct {
	NameType string `json:"nametype"`
}
