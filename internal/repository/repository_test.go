package repository

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock expected behavior
	mock.ExpectExec("INSERT INTO schema.table").
		WithArgs("aPPle", 5, "Fruit", "Local Market", sqlmock.AnyArg(), "John").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewPostgresqlDb(db, "schema", "table")

	item := models.Item{
		Name:     "aPPle",
		Quantity: 5,
		Type:     "Fruit",
		Market:   "Local Market",
		AddedAt:  time.Now(),
		AddedBy:  "John",
	}

	err = repo.CreateItem(item)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
