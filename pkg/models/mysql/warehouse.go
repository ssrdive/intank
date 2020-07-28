package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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

func (m *Warehouse) SecNumberModel(primaryNumber string) (models.SecNumberModel, error) {
	var secMod models.SecNumberModel

	err := m.DB.QueryRow(queries.SEC_MODEL, primaryNumber).Scan(&secMod.SecondaryNumber, &secMod.Model)

	if err != nil {
		return models.SecNumberModel{}, err
	}

	return secMod, nil
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

func (m *Warehouse) Movement(form url.Values) (int64, error) {
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

	var movementItems []models.GoodsMovement
	json.Unmarshal([]byte(form.Get("goods")), &movementItems)

	var primaryIDs []string

	primaryIDStr := ""

	for i, item := range movementItems {
		primaryIDs = append(primaryIDs, item.PrimaryNumber)
		primaryIDStr += item.PrimaryNumber
		if i != len(movementItems)-1 {
			primaryIDStr += ","
		}
	}

	var res []models.ValidTransfer
	err = mysequel.QueryToStructs(&res, m.DB, queries.INVALID_TRANSFERS(primaryIDStr), form.Get("warehouse_id"))
	if err != nil {
		return 0, err
	}

	fmt.Println(len(movementItems))
	fmt.Println(len(res))

	if len(movementItems) != len(res) {
		return 0, errors.New("Invalid Transfers")
	}

	for _, shentry := range res {
		_, err := mysequel.Insert(mysequel.Table{
			TableName: "stock_history",
			Columns:   []string{"document_id", "model_id", "primary_id", "secondary_id", "price", "date_in", "date_out"},
			Vals:      []interface{}{shentry.DocumentID, shentry.ModelID, shentry.PrimaryID, shentry.SecondaryID, shentry.Price, shentry.Date.Format("2006-01-02"), form.Get("date")},
			Tx:        tx,
		})
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	did, err := mysequel.Insert(mysequel.Table{
		TableName: "document",
		Columns:   []string{"document_type_id", "warehouse_id", "from_warehouse_id", "date"},
		Vals:      []interface{}{1, form.Get("warehouse_id"), form.Get("from_warehouse_id"), form.Get("date")},
		Tx:        tx,
	})
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(fmt.Sprintf("DELETE FROM main_stock WHERE primary_id IN (%s)", primaryIDStr))
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, shentry := range res {
		_, err := mysequel.Insert(mysequel.Table{
			TableName: "main_stock",
			Columns:   []string{"document_id", "model_id", "primary_id", "secondary_id", "price"},
			Vals:      []interface{}{did, shentry.ModelID, shentry.PrimaryID, shentry.SecondaryID, shentry.Price, form.Get("date")},
			Tx:        tx,
		})
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return did, nil
}

func (m *Warehouse) GoodsIn(form url.Values) (int64, error) {
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

	var goodsInItems []models.GoodsInItem
	json.Unmarshal([]byte(form.Get("goods")), &goodsInItems)

	id, err := mysequel.Insert(mysequel.Table{
		TableName: "document",
		Columns:   []string{"document_type_id", "warehouse_id", "from_warehouse_id", "date"},
		Vals:      []interface{}{1, form.Get("warehouse_id"), form.Get("from_warehouse_id"), form.Get("date")},
		Tx:        tx,
	})
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, item := range goodsInItems {
		_, err := mysequel.Insert(mysequel.Table{
			TableName: "main_stock",
			Columns:   []string{"document_id", "model_id", "primary_id", "secondary_id", "price"},
			Vals:      []interface{}{id, item.Model, item.PrimaryNumber, item.SecondaryNumber, item.Price},
			Tx:        tx,
		})
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, nil
}
