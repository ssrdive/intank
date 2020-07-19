package mysql

import (
	"database/sql"
	"net/url"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

// Warehouse struct holds methods to query item table
type Warehouse struct {
	DB *sql.DB
}

// Create creates an item
func (m *Warehouse) Create(rparams, oparams []string, form url.Values) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	id, err := mysequel.Insert(mysequel.FormTable{
		TableName: "warehouse",
		RCols:     rparams,
		OCols:     oparams,
		Form:      form,
		Tx:        tx,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// All returns all items
func (m *Warehouse) All() ([]models.AllWarehouseItem, error) {
	var res []models.AllWarehouseItem
	err := mysequel.QueryToStructs(&res, m.DB, queries.ALL_WAREHOUSES)
	if err != nil {
		return nil, err
	}

	return res, nil
}
