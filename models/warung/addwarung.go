package warung

type AddWarungs struct {
	OwnersID   int    `json:"ownersId"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	OwnersName string `json:"ownersName"`
}
