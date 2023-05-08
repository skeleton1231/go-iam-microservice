// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	srvv1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
)

// ItemController creates an item handler used to handle requests for the item resource.
type ItemController struct {
	srv srvv1.Service
}

// NewItemController creates an item handler.
func NewItemController(store store.Factory) *ItemController {
	return &ItemController{
		srv: srvv1.NewService(store),
	}
}
