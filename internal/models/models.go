package models

import "time"

// Item is the struct for the market item
type Item struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Quantity int       `db:"qty"`
	Type     string    `db:"type"`
	Market   string    `db:"market"`
	AddedAt  time.Time `db:"date"`
	AddedBy  string    `db:"added_by"`
}
