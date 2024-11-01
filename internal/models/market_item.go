package models

import (
	"github.com/pkg/errors"
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

func (i *Item) NormalizeFieldsForPersistence() error {
	var err error

	i.Name, err = normalizeMandatoryStrings(i.Name)
	if err != nil {
		return errors.Wrap(err, "unable to normalize 'name' field")
	}

	if i.Quantity <= 0 {
		return errors.Errorf("'quantity' must be greater than 0")
	}

	i.Type, err = normalizeMandatoryStrings(i.Type)
	if err != nil {
		return errors.Wrap(err, "unable to normalize 'type' field")
	}

	i.Market, err = normalizeMandatoryStrings(i.Market)
	if err != nil {
		return errors.Wrap(err, "unable to normalize 'market' field")
	}

	if i.AddedAt.IsZero() {
		return errors.Errorf("'added_at' timestamp is missing or not set")
	}

	i.AddedBy, err = normalizeMandatoryStrings(i.AddedBy)
	if err != nil {
		return errors.Wrap(err, "unable to normalize 'added_by' field")
	}

	return nil
}

func normalizeMandatoryStrings(str string) (string, error) {
	str = strings.ToLower(str)
	strs := strings.Fields(str)

	if len(strs) == 0 {
		return "", errors.Errorf("mandatory field can't be empty")
	}

	for idx, word := range strs {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		strs[idx] = string(runes)
	}

	str = strings.Join(strs, " ")

	return str, nil
}
