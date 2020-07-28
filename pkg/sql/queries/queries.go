package queries

import "fmt"

const ALL_MODELS = `
	SELECT * FROM model`

const ALL_WAREHOUSES = `
	SELECT W.id, WT.name as warehouse_type, W.name, W.address, W.contact FROM warehouse W LEFT JOIN warehouse_type WT ON WT.id = W.warehouse_type_id
`

const SEC_MODEL = `
	SELECT M.name, MS.secondary_id
	FROM main_stock MS
	LEFT JOIN model M ON M.id = MS.model_id
	WHERE primary_id = ?
`

func INVALID_TRANSFERS(pids string) string {
	return fmt.Sprintf("SELECT MS.*, DD.date FROM main_stock MS LEFT JOIN document DD ON MS.document_id = DD.id WHERE DD.warehouse_id = ? AND MS.primary_id IN (%s)", pids)
}
