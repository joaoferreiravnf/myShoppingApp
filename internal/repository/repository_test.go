package repository

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/mocks"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBExecutor(ctrl)

	repo := NewPostgresqlDb(mockDB, "public", "shopping_items")

	item := models.Item{
		Name:     "Apple",
		Quantity: 2,
		Type:     "Fruit",
		Market:   "Test Market",
		AddedAt:  time.Now(),
		AddedBy:  "Tester",
	}

	t.Run("successful item creation", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy).
			Return(nil, nil)

		err := repo.CreateItem(context.Background(), item)
		require.NoError(t, err, "expected no error when creating item")
	})
	t.Run("failed item creation due to db connection error", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy).
			Return(nil, sql.ErrConnDone)

		err := repo.CreateItem(context.Background(), item)
		assert.Contains(t, err.Error(), "connection is already closed")
	})
	t.Run("failed item creation due to context deadline exceeded", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy).
			Return(nil, context.DeadlineExceeded)

		err := repo.CreateItem(context.Background(), item)
		assert.Contains(t, err.Error(), "context deadline exceeded")
	})
}

func TestDeleteItem(t *testing.T) {

}
