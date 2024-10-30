package models

import (
	"strings"
	"time"
	"unicode"
)

// Item is the struct for the market item
type Item struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Quantity int       `db:"qty"`
	Type     string    `db:"type"`
	Market   string    `db:"market"`
	AddedAt  time.Time `db:"added_at"`
	AddedBy  string    `db:"added_by"`
}

func (i *Item) NormalizeNameForPersistence() {
	i.Name = strings.ReplaceAll(i.Name, "_", " ")
	i.Name = strings.ToLower(i.Name)

	uiName := strings.Fields(i.Name)

	for idx, word := range uiName {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		uiName[idx] = string(runes)
	}

	i.Name = strings.Join(uiName, " ")
}
