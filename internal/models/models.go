package models

import "time"

type MarketItem struct {
	Name     string    `db:"name"`
	Quantity int       `db:"qty"`
	Type     string    `db:"type"`
	Market   string    `db:"market"`
	AddedAt  time.Time `db:"date"`
	AddedBy  string    `db:"added_by"`
}
