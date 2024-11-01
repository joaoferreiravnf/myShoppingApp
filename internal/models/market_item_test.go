package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeFieldsForPersistence(t *testing.T) {
	item := Item{}
	testCases := []struct {
		UiName string
		dbName string
	}{
		{"Red markers", "red_markers"},
		{"markers", "markers"},
		{"", ""},
		{"marKERS", "markers"},
		{" marKers", "markers"},
	}

	for _, tc := range testCases {
		item.Name = tc.UiName
		err := item.NormalizeFieldsForPersistence()
		assert.NoError(t, err)
		assert.Equal(t, tc.dbName, item.Name)
	}
}

func TestNormalizeNameForPersistence(t *testing.T) {
	item := Item{}
	testCases := []struct {
		uiName string
		dbName string
	}{
		{"Red markers", "Red Markers"},
		{"markers", "Markers"},
		{"", ""},
		{"marKERS", "Markers"},
		{" marKers", "Markers"},
	}

	for _, tc := range testCases {
		item.Name = tc.dbName
		item.NormalizeFieldsForPersistence()
		assert.Equal(t, tc.dbName, item.Name)
	}
}
