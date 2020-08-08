package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type User struct {
	ID        int
	GroupID   int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
}

type JWTUser struct {
	ID       int
	Username string
	Password string
	Name     string
	Type     string
}

type Dropdown struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AllItemItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Country       string `json:"country"`
	PrimaryName   string `json:"primary_name"`
	SecondaryName string `json:"secondary_name"`
}

type ItemDetails struct {
	ID               int     `json:"id"`
	ItemID           string  `json:"item_id"`
	ModelID          string  `json:"model_id"`
	ModelName        string  `json:"model_name"`
	ItemCategoryID   string  `json:"item_category_id"`
	ItemCategoryName string  `json:"item_category_name"`
	PageNo           string  `json:"page_no"`
	ItemNo           string  `json:"item_no"`
	ForeignID        string  `json:"foreign_id"`
	ItemName         string  `json:"item_name"`
	Price            float64 `json:"price"`
}

type AllWarehouseItem struct {
	ID            int    `json:"id"`
	WarehouseType string `json:"warehouse_type"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Contact       string `json:"contact"`
}

type GoodsInItem struct {
	Model           string
	PrimaryNumber   string `json:"primary_number"`
	SecondaryNumber string `json:"secondary_number"`
	Price           string
}

type GoodsMovement struct {
	Model           string `json:"model"`
	PrimaryNumber   string `json:"primaryNumber"`
	SecondaryNumber string `json:"secondaryNumber"`
	Price           string `json:"price"`
}

type SecNumberModel struct {
	SecondaryNumber string `json:"secondaryNumber"`
	Model           string `json:"model"`
}

type ValidTransfer struct {
	DocumentID  string
	ModelID     string
	PrimaryID   string
	SecondaryID string
	Price       string
	Date        time.Time
}

type WarehouseStockItem struct {
	DocumentID           int    `json:"document_id"`
	PrimaryID            string `json:"primary_id"`
	SecondaryID          string `json:"secondary_id"`
	InStockFor           int    `json:"in_stock_for"`
	Price                int    `json:"price"`
	Model                string `json:"model"`
	Date                 string `json:"date"`
	DeliveryDocumentType string `json:"delivery_document_type"`
}

type SearchResultItem struct {
	DocumentID  int    `json:"document_id"`
	Model       string `json:"model"`
	Warehouse   string `json:"warehouse"`
	PrimaryID   string `json:"primary_id"`
	SecondaryID string `json:"secondary_id"`
	Price       int    `json:"price"`
	WarehouseID int    `json:"warehouse_id"`
}

type AgeWiseItem struct {
	DocumentID           int    `json:"document_id"`
	PrimaryID            string `json:"primary_id"`
	SecondaryID          string `json:"secondary_id"`
	InStockFor           int    `json:"in_stock_for"`
	Price                int    `json:"price"`
	Model                string `json:"model"`
	Date                 string `json:"date"`
	DeliveryDocumentType string `json:"delivery_document_type"`
}

type HistoryItem struct {
	DocumentID           int    `json:"document_id"`
	PrimaryID            string `json:"primary_id"`
	SecondaryID          string `json:"secondary_id"`
	InStockFor           int    `json:"in_stock_for"`
	Price                int    `json:"price"`
	DateIn               string `json:"date_in"`
	DateOut              string `json:"date_out"`
	DeliveryDocumentType string `json:"delivery_document_type"`
	Warehouse            string `json:"warehouse"`
	WarehouseID          int    `json:"warehouse_id"`
	Model                string `json:"model"`
}

type DocsItem struct {
	DocumentID           int    `json:"document_id"`
	DeliveryDocumentType string `json:"delivery_document_type"`
	Date                 string `json:"date"`
	ToWarehouseID        string `json:"to_warehouse_id"`
	ToWarehouse          string `json:"to_warehouse"`
	FromWarehouseID      string `json:"from_warehouse_id"`
	FromWarehouse        string `json:"from_warehouse"`
}

type StockByModel struct {
	Model string `json:"model"`
	Count int    `json:"count"`
}
