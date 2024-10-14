package repository

import (
	"database/sql"
	"fmt"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
)

type Repository interface {
	CreateItem(item models.Item) error
	ListItems() ([]models.Item, error)
	DeleteItem(name string) error
	EditItem(item models.Item) error
}

// PostgresqlDb is the current db concrete implementation
type PostgresqlDb struct {
	db      *sql.DB
	dbName  string
	dbTable string
}

// NewPostgresqlDb creates a new PostgresqlDb instance
func NewPostgresqlDb(db *sql.DB, dbName, dbTable string) *PostgresqlDb {
	return &PostgresqlDb{
		db:      db,
		dbName:  dbName,
		dbTable: dbTable,
	}
}

// CreateItem creates a new item in the database
func (ir *PostgresqlDb) CreateItem(item models.Item) error {
	query := fmt.Sprintf("INSERT INTO %s.%s (name, qty, type, market, date, added_by) VALUES ($1, $2, $3, $4, $5, $6)", ir.dbName, ir.dbTable)

	_, err := ir.db.Exec(query, item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy)

	if err != nil {
		return err
	}

	return nil
}

func (ir *PostgresqlDb) ListItems() ([]models.Item, error) {
	rows, err := ir.db.Query(fmt.Sprintf("SELECT * FROM %s.%s"), ir.dbName, ir.dbTable)
	if err != nil {
		return nil, err
	}

	var items []models.Item

	for rows.Next() {
		var item models.Item

		err = rows.Scan(&item.Name, &item.Quantity, &item.Type, &item.Market, &item.AddedAt, &item.AddedBy)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (ir PostgresqlDb) EditItem(item models.Item) error {
	return nil
}

func (ir PostgresqlDb) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s.%s WHERE id=$1", ir.dbName, ir.dbTable)

	_, err := ir.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
