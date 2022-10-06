package model

import "time"

type Company struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	OIB       int64     `json:"oib"`
	Address   string    `json:"address"`
	Deleted   bool      `json:"-"`
	DeletedAt time.Time `json:"-"`
}
