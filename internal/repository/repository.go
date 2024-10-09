package repository

import (
	"database/sql"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
)

type Repository interface {
	CreateItem(item models.MarketItem) error
}

// PostgresqlDb is the current db concrete implementation
type PostgresqlDb struct {
	db *sql.DB
}

// NewPostgresqlDb creates a new PostgresqlDb instance
func NewPostgresqlDb(db *sql.DB) *PostgresqlDb {
	return &PostgresqlDb{db: db}
}

// CreateItem creates a new item in the database
func (ir *PostgresqlDb) CreateItem(item models.MarketItem) error {
	_, err := ir.db.Exec("INSERT INTO shopping_app.shopping_list(name, qty, type, market, date, added_by) "+
		"VALUES($1, $2, $3, $4, $5, $6)",
		item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy)

	if err != nil {
		return err
	}

	err = ir.db.Close()

	if err != nil {
		return err
	}

	return nil
}

func (ir *PostgresqlDb) ListItems() ([]models.MarketItem, error) {
	rows, err := ir.db.Query("SELECT * FROM shopping_app.shopping_list")
	if err != nil {
		return nil, err
	}

	var items []models.MarketItem

	for rows.Next() {
		var item models.MarketItem

		err := rows.Scan(&item.Name, &item.Quantity, &item.Type, &item.Market, &item.AddedAt, &item.AddedBy)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	err = ir.db.Close()

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (ir PostgresqlDb) DeleteItem(name string) error {
	_, err := ir.db.Exec("DELETE FROM shopping_app.shopping_list WHERE name=$1", name)

	if err != nil {
		return err
	}

	err = ir.db.Close()

	if err != nil {
		return err
	}

	return nil
}
