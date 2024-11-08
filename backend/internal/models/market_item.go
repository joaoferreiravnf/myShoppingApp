package models

import (
	"github.com/pkg/errors"
	"slices"
	"strings"
	"time"
	"unicode"
)

const (
	Continente = "Continente"
	Lidl       = "Lidl"
	Belita     = "Belita"
	PingoDoce  = "Pingo Doce"
	Jumbo      = "Jumbo"
)

const (
	Fruits      = "Frutas"
	Vegetables  = "Legumes"
	MeatAndFish = "Carne & Peixe"
	Drinks      = "Bebidas"
	Higiene     = "Higiene"
)

// TODO: Consider imterface implementation for future structs (multiuse app)

// Item is the struct for the market item
type Item struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Quantity int    `db:"qty"`
	Type     string `db:"type"`
	Market   string `db:"market"`
	AddedAt  string `db:"added_at"`
	AddedBy  string `db:"added_by"`
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

	i.AddedAt = time.Now().Format("2006-01-02")

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

type ListItemsData struct {
	Items   []Item
	Markets []string
	Types   []string
}

func (lid *ListItemsData) GetMarkets() {
	markets := []string{
		Continente,
		Belita,
		PingoDoce,
		Jumbo,
		Lidl,
	}

	slices.Sort(markets)

	lid.Markets = markets
}

func (lid *ListItemsData) GetTypes() {
	types := []string{
		Fruits,
		Vegetables,
		MeatAndFish,
		Drinks,
		Higiene,
	}

	lid.Types = types
}
