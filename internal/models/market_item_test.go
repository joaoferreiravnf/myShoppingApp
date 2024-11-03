package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNormalizeFieldsForPersistence(t *testing.T) {
	t.Run("successful fields normalization", func(t *testing.T) {
		items := []Item{
			{Name: " aPple ", Quantity: 10, Type: " frUitS ", Market: " Test market ", AddedAt: time.Now(), AddedBy: " John "},
		}

		for _, item := range items {
			err := item.NormalizeFieldsForPersistence()
			assert.NoError(t, err)
			assert.Equal(t, "Apple", item.Name)
			assert.Equal(t, 10, item.Quantity)
			assert.Equal(t, "Fruits", item.Type)
			assert.Equal(t, "Test Market", item.Market)
			assert.Equal(t, "John", item.AddedBy)
		}
	})
	t.Run("failed fields normalization due to empty fields", func(t *testing.T) {
		item := Item{}

		err := item.NormalizeFieldsForPersistence()
		assert.ErrorContains(t, err, "mandatory field can't be empty")
	})
	t.Run("failed fields normalization due to less than 0 quantity", func(t *testing.T) {
		item := Item{Name: " aPple ", Quantity: 0, Type: " frUitS ", Market: " Test market ", AddedAt: time.Now(), AddedBy: " John "}

		err := item.NormalizeFieldsForPersistence()
		assert.ErrorContains(t, err, "'quantity' must be greater than 0")
	})
}
