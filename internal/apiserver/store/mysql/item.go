package mysql

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"gorm.io/gorm"
)

type items struct {
	db *gorm.DB
}

func newItems(ds *datastore) *items {
	return &items{ds.db}
}

// Create creates a new item.
func (i *items) Create(ctx context.Context, item *v1.Item, opts metav1.CreateOptions) error {
	return i.db.Create(&item).Error
}

// Update updates an item.
func (i *items) Update(ctx context.Context, item *v1.Item, opts metav1.UpdateOptions) error {
	return i.db.Save(item).Error
}

// Delete deletes an item by the item ID.
func (i *items) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		i.db = i.db.Unscoped()
	}

	err := i.db.Where("id = ?", id).Delete(&v1.Item{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Get returns an item by the item identifier.
func (i *items) Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.Item, error) {
	item := &v1.Item{}
	err := i.db.Where("id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUnknown, err.Error()) // code.ErrItemNotFound
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return item, nil
}

// List returns all items.
// func (i *items) List(ctx context.Context, opts metav1.ListOptions) ([]*v1.ItemList, error) {
// 	ret := []*v1.ItemList{}
// 	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

// 	d := i.db.Offset(ol.Offset).
// 		Limit(ol.Limit).
// 		Order("id desc").
// 		Preload("ItemAttributes").
// 		Preload("ItemImage").
// 		Preload("ItemSummaryByMarketplace").
// 		Preload("Issue").
// 		Preload("ItemOfferByMarketplace").
// 		Preload("ItemProcurement").
// 		Find(&ret).
// 		Offset(-1).
// 		Limit(-1).
// 		Count()

// 	return ret, d.Error
// }
