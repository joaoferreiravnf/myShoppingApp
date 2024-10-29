package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeStringForPersistence(t *testing.T) {
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
		item.NormalizeNameForPersistence()
		assert.Equal(t, tc.dbName, item.Name)
	}
}

func TestNormalizeStringForUI(t *testing.T) {
	item := Item{}
	testCases := []struct {
		dbName string
		UiName string
	}{
		{"red_markers", "Red Markers"},
		{"markers", "Markers"},
		{"", ""},
		{"_blaCK_marKErs_", "Black Markers"},
	}

	for _, tc := range testCases {
		item.Name = tc.UiName
		item.NormalizeNameForUI()
		assert.Equal(t, tc.UiName, item.Name)
	}
}
