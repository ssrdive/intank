package queries

const ALL_MODELS = `
	SELECT * FROM model`

const ALL_WAREHOUSES = `
	SELECT W.id, WT.name as warehouse_type, W.name, W.address, W.contact FROM warehouse W LEFT JOIN warehouse_type WT ON WT.id = W.warehouse_type_id
`
