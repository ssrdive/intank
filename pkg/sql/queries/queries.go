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

const WAREHOUSE_STOCK = `
	SELECT MS.document_id, MS.primary_id, MS.secondary_id, DATEDIFF(NOW(), DD.date) as in_stock_for, MS.price, M.name as model, DD.date, DDT.name as delivery_document_type 
	FROM main_stock MS 
	LEFT JOIN model M  ON MS.model_id = M.id 
	LEFT JOIN document DD ON MS.document_id = DD.id 
	LEFT JOIN document_type DDT ON DD.document_type_id = DDT.id 
	WHERE DD.warehouse_id = ?
`

const SEARCH = `
	SELECT MS.document_id, M.name AS model, W.name as warehouse, MS.primary_id, MS.secondary_id, MS.price, W.id as warehouse_id 
	FROM main_stock MS 
	LEFT JOIN model M ON MS.model_id = M.id 
	LEFT JOIN document DD ON MS.document_id = DD.id 
	LEFT JOIN warehouse W ON DD.warehouse_id = W.id 
	WHERE CONCAT(MS.document_id, M.name, W.name, MS.primary_id, MS.secondary_id) LIKE ?
`

const AGE_WISE_SEARCH = `
	SELECT MS.document_id, MS.primary_id, MS.secondary_id, DATEDIFF(NOW(), DD.date) as in_stock_for, MS.price, M.name as model, DD.date, DDT.name as delivery_document_type 
	FROM main_stock MS 
	LEFT JOIN model M  ON MS.model_id = M.id 
	LEFT JOIN document DD ON MS.document_id = DD.id 
	LEFT JOIN document_type DDT ON DD.document_type_id = DDT.id 
	WHERE DATEDIFF(NOW(), DD.date) >= ? AND MS.model_id = ?
`
