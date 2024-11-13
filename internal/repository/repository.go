package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/pkg/errors"
)

type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// Repository defines the methods for managing items in a persistence state
type Repository interface {
	CreateItem(ctx context.Context, item models.Item) error
	ListItems(ctx context.Context) ([]models.Item, error)
	UpdateItem(ctx context.Context, item models.Item) error
	DeleteItem(ctx context.Context, item models.Item) error
}

// PostgresqlDb is the current database concrete implementation
type PostgresqlDb struct {
	db       DBExecutor
	dbSchema string
	dbTable  string
}

// NewPostgresqlDb creates a new PostgresqlDb instance
func NewPostgresqlDb(db DBExecutor, dbSchema, dbTable string) *PostgresqlDb {
	return &PostgresqlDb{
		db:       db,
		dbSchema: dbSchema,
		dbTable:  dbTable,
	}
}

// CreateItem creates a new item in the table
func (ir *PostgresqlDb) CreateItem(ctx context.Context, item models.Item) error {
	query := fmt.Sprintf("INSERT INTO %s.%s (name, qty, type, market, added_at, added_by) VALUES ($1, $2, $3, $4, $5, $6)", ir.dbSchema, ir.dbTable)

	_, err := ir.db.ExecContext(ctx, query, item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy)
	if err != nil {
		return errors.Wrap(err, "error inserting new item into the database")
	}

	return nil
}

// ListItems lists all the items of the table
func (ir *PostgresqlDb) ListItems(ctx context.Context) ([]models.Item, error) {
	rows, err := ir.db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s.%s ORDER BY market ASC", ir.dbSchema, ir.dbTable))
	if err != nil {
		return nil, errors.Wrap(err, "error querying items from database")
	}

	var items []models.Item

	for rows.Next() {
		var item models.Item

		err = rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Type, &item.Market, &item.AddedAt, &item.AddedBy)
		if err != nil {
			return nil, errors.Wrap(err, "error converting values from table rows into item object")
		}

		items = append(items, item)
	}

	return items, nil
}

func (ir PostgresqlDb) UpdateItem(ctx context.Context, item models.Item) error {
	query := fmt.Sprintf(`UPDATE %s.%s SET name = $1, qty = $2, type = $3, market = $4, added_at = $5, 
    	added_by = $6 WHERE id = $7
		`, ir.dbSchema, ir.dbTable)

	_, err := ir.db.ExecContext(ctx, query, item.Name, item.Quantity, item.Type, item.Market, item.AddedAt, item.AddedBy, item.ID)
	if err != nil {
		return errors.Wrapf(err, "error updating item %s with id %d", item.Name, item.ID)
	}

	return nil
}

// DeleteItem deletes the item from the table
func (ir PostgresqlDb) DeleteItem(ctx context.Context, item models.Item) error {
	query := fmt.Sprintf("DELETE FROM %s.%s WHERE id=$1", ir.dbSchema, ir.dbTable)

	_, err := ir.db.ExecContext(ctx, query, item.ID)
	if err != nil {
		return errors.Wrapf(err, "error deleting item %s with id %d from database", item.Name, item.ID)
	}

	return nil
}
