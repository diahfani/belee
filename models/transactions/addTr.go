package transactions

type Addtransactions struct {
	ProductsName string  `json:"productsname"`
	TotalQty     int     `json:"totalqty"`
	Totalprice   float32 `json:"totalprice"`
	// Status       string  `json:"status"`
}
