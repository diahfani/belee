package warung

import "time"

type AddWarungs struct {
	OwnersID int    `json:"ownersId"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	// OwnersName string `json:"ownersName"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
