package models

import "time"

// MarketItem is the struct for the market item
type MarketItem struct {
	Name     string    `db:"name"`
	Quantity int       `db:"qty"`
	Type     string    `db:"type"`
	Market   string    `db:"market"`
	AddedAt  time.Time `db:"date"`
	AddedBy  string    `db:"added_by"`
}
