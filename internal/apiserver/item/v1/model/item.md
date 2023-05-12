### Q&A

### bugs
1. 2023-05-12 13:23:10.029 INFO    apiserver       gorm@v1.22.4/callbacks.go:134   [mysql/item.go:33][42.875ms] [rows:1] UPDATE `item` SET `asin`='B0B3HXK9GS',`sku`='LU-US-01',`brand`='Kuggini',`title`='Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX7 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 1000',`product_group`='',`product_type`='',`created_at`=NULL,`updated_at`='2023-05-12 13:23:09.988' WHERE `id` = 5

It seems that the issue is with the `created_at` column being set to `'0000-00-00 00:00:00'`, which is not a valid datetime value. To resolve this, you can modify the `created_at` column in your database schema to allow `NULL` values and update the `Item` struct accordingly.

Here's the updated code:

```go
type Item struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	ASIN         string    `json:"asin"`
	SKU          string    `json:"sku"`
	Brand        string    `json:"brand"`
	Title        string    `json:"title"`
	ProductGroup string    `json:"product_group"`
	ProductType  string    `json:"product_type"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
```

By changing the `CreatedAt` field type to `*time.Time`, it allows the value to be `NULL` in the database. GORM will automatically handle the `NULL` value for the `created_at` column during insertion.

Make sure to update your database schema accordingly to allow `NULL` values for the `created_at` column.
