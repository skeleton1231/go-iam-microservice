// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	"net/http"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

// Create creates a new item.
func (ic *ItemController) Create(c *gin.Context) {
	var item v1.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrEncrypt, err.Error()), nil)

		return
	}

	id := node.Generate()
	item.ID = uint64(id.Int64())
	idStr := strconv.Itoa(int(item.ID))
	item.InstanceID = "item-" + idStr
	name := "sku-" + idStr
	item.Name = name
	item.SKU = name

	if err := ic.srv.Items().Create(c, &item, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)

		return
	}

	c.JSON(http.StatusCreated, item)
}
