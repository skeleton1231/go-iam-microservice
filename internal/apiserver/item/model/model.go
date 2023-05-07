package model

import "time"

type Item struct {
	ID           int       `json:"id"`
	ASIN         string    `json:"asin"`
	SKU          string    `json:"sku"`
	Brand        string    `json:"brand"`
	Title        string    `json:"title"`
	ProductGroup string    `json:"product_group"`
	ProductType  string    `json:"product_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ItemAttributes struct {
	ID                    int       `json:"id"`
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
	ID        int       `json:"id"`
	ItemID    int       `json:"item_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemSummaryByMarketplace struct {
	ID            int       `json:"id"`
	ItemID        int       `json:"item_id"`
	MarketplaceID string    `json:"marketplace_id"`
	SalesRank     int       `json:"sales_rank"`
	MainImageURL  string    `json:"main_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Issue struct {
	ID        int       `json:"id"`
	ItemID    int       `json:"item_id"`
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Severity  string    `json:"severity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemOfferByMarketplace struct {
	ID                 int       `json:"id"`
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
	ID                    int       `json:"id"`
	ItemID                int       `json:"item_id"`
	ExternalProductID     string    `json:"external_product_id"`
	ExternalProductIDType string    `json:"external_product_id_type"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type FulfillmentAvailability struct {
	ID                  int       `json:"id"`
	OfferID             int       `json:"offer_id"`
	FulfillmentCenterID string    `json:"fulfillment_center_id"`
	QuantityAvailable   int       `json:"quantity_available"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
