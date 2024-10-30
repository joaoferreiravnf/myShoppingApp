package repository

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/mocks"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock DBExecutor
	mockDB := mocks.NewMockDBExecutor(ctrl)

	// Define repository with mock
	repo := NewPostgresqlDb(mockDB, "public", "shopping_items")

	item := models.Item{
		Name:     "Test Item",
		Quantity: 3,
		Type:     "Grocery",
		Market:   "Test Market",
		AddedAt:  time.Now(),
		AddedBy:  "Tester",
	}

	// Set expectation for ExecContext on the mock
	mockDB.EXPECT().
		ExecContext(gomock.Any(), gomock.Any(), item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy).
		Return(nil, nil) // Mock the successful return

	// Call CreateItem and verify
	err := repo.CreateItem(context.Background(), item)
	require.NoError(t, err, "expected no error when creating item")
}
