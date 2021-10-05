package transactions

type Addtransactions struct {
	BuyerId      int     `json:"buyerId"`
	WarungId     int     `json:"warungId"`
	BarangId     int     `json:"barangId"`
	PaymentId    int     `json:"paymentId"`
	ProductsName string  `json:"productsname"`
	TotalQty     int     `json:"totalqty"`
	Totalprice   float32 `json:"totalprice"`
	// Status       string  `json:"status"`
}
