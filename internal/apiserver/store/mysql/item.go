// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"fmt"

	"github.com/marmotedu/component-base/pkg/fields"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/util/gormutil"
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

func (i *items) List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error) {
	ret := &v1.ItemList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	selector, _ := fields.ParseSelector(opts.FieldSelector)

	// Build the query
	query := i.db.Where("status = 1")

	for _, req := range selector.Requirements() {
		switch req.Field {
		case "asin":
			query = query.Where("asin = ?", req.Value)
		case "sku":
			query = query.Where("sku = ?", req.Value)
		case "brand":
			query = query.Where("brand = ?", req.Value)
		case "title":
			query = query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", req.Value))
		case "product_group":
			query = query.Where("product_group = ?", req.Value)
		case "product_type":
			query = query.Where("product_type = ?", req.Value)
		}
	}

	// Execute the query
	d := query.
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
