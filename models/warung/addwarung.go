package warung

type AddWarungs struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	OwnersID   int    `json:"ownersId"`
	OwnersName string `json:"ownersName"`
}
