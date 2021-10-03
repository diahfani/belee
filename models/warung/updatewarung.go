package warung

import "time"

type UpdateWarung struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	UpdatedAt time.Time `json:"updatedAt"`
}
