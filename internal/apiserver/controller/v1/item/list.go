// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package item

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

func (i *ItemController) List(c *gin.Context) {
	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	items, err := i.srv.Items().List(c, r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, items)
}

// func (i *ItemController) getFilteredItems(opts options.FilteredListOptions) (*v1.ItemList, error) {
// 	items, err := i.getAllItems(opts.ListOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filteredItems := &v1.ItemList{Items: []*v1.Item{}}
// 	for _, item := range items.Items {
// 		if i.shouldIncludeItem(item, opts.Filter) {
// 			filteredItems.Items = append(filteredItems.Items, item)
// 		}
// 	}

// 	return filteredItems, nil
// }

// func (i *ItemController) shouldIncludeItem(item *v1.Item, filter string) bool {
// 	if filter == "" {
// 		return true
// 	}

// 	return false
// }

// func (i *ItemController) getAllItems(opts metav1.ListOptions) (*v1.ItemList, error) {
// 	itemList, err := i.srv.Items().List(context.Background(), opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return itemList, nil
// }
