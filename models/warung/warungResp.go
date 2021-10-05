package warung

import (
	"belee/models/owner"
	"time"
)

type WarungRes struct {
	Owner     *owner.Owners `json:"owner"`
	Name      string        `json:"name"`
	Address   string        `json:"address"`
	CreatedAt time.Time     `json:"createdat"`
	UpdatedAt time.Time     `json:"updatedAt"`
}
