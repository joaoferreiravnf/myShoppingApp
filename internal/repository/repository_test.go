package repository

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/mocks"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPostgresqlDb_CreateItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBExecutor(ctrl)

	repo := NewPostgresqlDb(mockDB, "public", "shopping_items")

	item := models.Item{
		Name:     "Apple",
		Quantity: 2,
		Type:     "Fruit",
		Market:   "Test Market",
		AddedAt:  time.Now().Format("02-01"),
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
	t.Run("failed item creation due to duplicate item insertion", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy).
			Return(nil, errors.Errorf("duplicate key value violates unique constraint"))

		err := repo.CreateItem(context.Background(), item)
		assert.Contains(t, err.Error(), "duplicate key value violates unique constraint")
	})
}

func TestPostgresqlDb_DeleteItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBExecutor(ctrl)
	ctx := context.Background()
	repo := NewPostgresqlDb(mockDB, "public", "shopping_items")

	item := models.Item{
		ID:       3,
		Name:     "Apple",
		Quantity: 2,
		Type:     "Fruit",
		Market:   "Test Market",
		AddedAt:  time.Now().Format("02-01"),
		AddedBy:  "Tester",
	}

	t.Run("successful item deletion", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.ID).
			Return(nil, nil)

		err := repo.DeleteItem(ctx, item)
		assert.NoError(t, err)
	})
	t.Run("failed item deletion due to db connection error", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.ID).
			Return(nil, sql.ErrConnDone)

		err := repo.DeleteItem(context.Background(), item)
		assert.Contains(t, err.Error(), "connection is already closed")
	})
	t.Run("failed item deletion due to context deadline exceeded", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.ID).
			Return(nil, context.DeadlineExceeded)

		err := repo.DeleteItem(context.Background(), item)
		assert.Contains(t, err.Error(), "context deadline exceeded")
	})
	t.Run("failed item deletion due to non existent item", func(t *testing.T) {
		mockDB.EXPECT().
			ExecContext(gomock.Any(), gomock.Any(), item.ID).
			Return(nil, sql.ErrNoRows)

		err := repo.DeleteItem(context.Background(), item)
		assert.Contains(t, err.Error(), "no rows in result set")
	})
}

func TestPostgresqlDb_UpdateItem(t *testing.T) {

}
