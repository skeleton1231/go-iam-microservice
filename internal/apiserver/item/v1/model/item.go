// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"time"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type ItemWithDetails struct {
	Item                     Item
	ItemAttributes           ItemAttributes
	ItemImage                ItemImage
	ItemSummaryByMarketplace ItemSummaryByMarketplace
	Issue                    Issue
	ItemOfferByMarketplace   ItemOfferByMarketplace
	ItemProcurement          ItemProcurement
}

type ItemList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`

	Items []*Item `json:"items"`
}

//	type Item struct {
//		ID           int    `gorm:"primaryKey" json:"id"`
//		ASIN         string `json:"asin"`
//		SKU          string `json:"sku"`
//		Brand        string `json:"brand"`
//		Title        string `json:"title"`
//		ProductGroup string `json:"product_group"`
//		ProductType  string `json:"product_type"`
//		// CreatedAt    time.Time `json:"created_at"`
//		// UpdatedAt    time.Time `json:"updated_at"`
//		CreatedAt time.Time `gorm:"-" json:"created_at"`
//		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
//	}
type Item struct {
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ASIN         string `json:"asin"`
	SKU          string `json:"sku"`
	Brand        string `json:"brand"`
	Title        string `json:"title"`
	ProductGroup string `json:"product_group"`
	ProductType  string `json:"product_type"`
}

func (Item) TableName() string {
	return "item" // specify the desired table name here
}

type ItemAttributes struct {
	ID                    int       `gorm:"primaryKey" json:"id"`
	ItemID                int       `json:"item_id"`
	Binding               string    `json:"binding"`
	ItemHeight            float64   `json:"item_height"`
	ItemLength            float64   `json:"item_length"`
	ItemWidth             float64   `json:"item_width"`
	ItemWeight            float64   `json:"item_weight"`
	ItemDimensionsUnit    string    `json:"item_dimensions_unit"`
	PackageHeight         float64   `json:"package_height"`
	PackageLength         float64   `json:"package_length"`
	PackageWidth          float64   `json:"package_width"`
	PackageWeight         float64   `json:"package_weight"`
	PackageDimensionsUnit string    `json:"package_dimensions_unit"`
	ReleaseDate           time.Time `json:"release_date"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type ItemImage struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	ItemID    int       `json:"item_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemSummaryByMarketplace struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	ItemID        int       `json:"item_id"`
	MarketplaceID string    `json:"marketplace_id"`
	SalesRank     int       `json:"sales_rank"`
	MainImageURL  string    `json:"main_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Issue struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	ItemID    int       `json:"item_id"`
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Severity  string    `json:"severity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemOfferByMarketplace struct {
	ID                 int       `gorm:"primaryKey" json:"id"`
	ItemID             int       `json:"item_id"`
	MarketplaceID      string    `json:"marketplace_id"`
	ListPrice          float64   `json:"list_price"`
	CurrencyCode       string    `json:"currency_code"`
	PackageQuantity    int       `json:"package_quantity"`
	AvailabilityStatus string    `json:"availability_status"`
	FulfillmentChannel string    `json:"fulfillment_channel"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ItemProcurement struct {
	ID                    int       `gorm:"primaryKey" json:"id"`
	ItemID                int       `json:"item_id"`
	ExternalProductID     string    `json:"external_product_id"`
	ExternalProductIDType string    `json:"external_product_id_type"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
