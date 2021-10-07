package transactions

//dihapus aja file ini
import "belee/models/products"

type BuyerTransactionResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DetailResponse struct {
	Id       int                      `json:"id"`
	Products products.Products        `json:"products"`
	Buyer    BuyerTransactionResponse `json:"buyer"`
}
