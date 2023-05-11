// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package item

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/options"
)

func (i *ItemController) List(c *gin.Context) {
	var opts options.FilteredListOptions
	if err := c.ShouldBindQuery(&opts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	items, err := i.getFilteredItems(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (i *ItemController) getFilteredItems(opts options.FilteredListOptions) (*v1.ItemList, error) {
	items, err := i.getAllItems(opts.ListOptions)
	if err != nil {
		return nil, err
	}

	filteredItems := &v1.ItemList{Items: []*v1.Item{}}
	for _, item := range items.Items {
		if i.shouldIncludeItem(item, opts.Filter) {
			filteredItems.Items = append(filteredItems.Items, item)
		}
	}

	return filteredItems, nil
}

func (i *ItemController) shouldIncludeItem(item *v1.Item, filter string) bool {
	if filter == "" {
		return true
	}

	if strings.Contains(strings.ToLower(item.ASIN), strings.ToLower(filter)) {
		return true
	}

	return false
}

func (i *ItemController) getAllItems(opts metav1.ListOptions) (*v1.ItemList, error) {
	itemList, err := i.srv.Items().List(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
