package mysql

import (
	"database/sql"
	"net/url"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

// MModel struct holds methods to query item table
type MModel struct {
	DB *sql.DB
}

// Create creates an item
func (m *MModel) Create(rparams, oparams []string, form url.Values) (int64, error) {
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
		TableName: "model",
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
func (m *MModel) All() ([]models.AllItemItem, error) {
	var res []models.AllItemItem
	err := mysequel.QueryToStructs(&res, m.DB, queries.ALL_MODELS)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MModel) UserAll() ([]models.AllUserItem, error) {
	var res []models.AllUserItem
	err := mysequel.QueryToStructs(&res, m.DB, queries.ALL_USERS)
	if err != nil {
		return nil, err
	}

	return res, nil
}
