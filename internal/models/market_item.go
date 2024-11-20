package models

import (
	"github.com/pkg/errors"
	"slices"
	"strings"
	"time"
	"unicode"
)

const (
	continente = "Continente"
	lidl       = "Lidl"
	belita     = "Belita"
	pingoDoce  = "Pingo Doce"
	jumbo      = "Jumbo"
)

const (
	fruits      = "Frutas"
	vegetables  = "Legumes"
	meatAndFish = "Carne & Peixe"
	drinks      = "Bebidas"
	higiene     = "Higiene"
)

const itemQuantities = 10

// TODO: Consider imterface implementation for future structs (multiuse app)

// Item is the struct for the market item
type Item struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Quantity int    `db:"qty" json:"qty"`
	Type     string `db:"type" json:"type"`
	Market   string `db:"market" json:"market"`
	AddedAt  string `db:"added_at" json:"added_at"`
	AddedBy  string `db:"added_by" json:"added_by"`
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

	i.AddedAt = time.Now().Format("02-01")

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
	Items      []Item   `json:"item"`
	Markets    []string `json:"markets"`
	Types      []string `json:"types"`
	Quantities []int    `json:"quantities"`
}

func (lid *ListItemsData) GetMarkets() {
	markets := []string{
		continente,
		belita,
		pingoDoce,
		jumbo,
		lidl,
	}

	slices.Sort(markets)

	lid.Markets = markets
}

func (lid *ListItemsData) GetTypes() {
	types := []string{
		fruits,
		vegetables,
		meatAndFish,
		drinks,
		higiene,
	}

	slices.Sort(types)

	lid.Types = types
}

func (lid *ListItemsData) GetQuantities() {
	var quantities []int

	for i := 1; i <= itemQuantities; i++ {
		quantities = append(quantities, i)
	}

	lid.Quantities = quantities
}
