// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/log"
)

// Update updates an item by its ID.
func (ic *ItemController) Update(c *gin.Context) {
	log.L(c).Info("update item function called.")

	itemID, _ := strconv.Atoi(c.Param("itemID"))

	var newItem v1.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	item, err := ic.srv.Items().Get(c, itemID, metav1.GetOptions{})

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	newItem.ID = item.ID

	// Save changed fields.
	if err := ic.srv.Items().Update(c, &newItem, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, newItem)

}
