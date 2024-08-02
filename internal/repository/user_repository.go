package repository

import (
	"database/sql"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(item models.MarketItem) error {
	_, err := ur.db.Exec("INSERT INTO shopping_app.shopping_list(name, qty, type, market, date, added_by) "+
		"VALUES($1, $2, $3, $4, $5, $6)",
		item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy)

	if err != nil {
		return err
	}
	return nil
}
