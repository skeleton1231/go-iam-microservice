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
	srvv1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

// ItemAttributesController creates an item attributes handler used to handle requests for the item attributes resource.
type ItemAttributesController struct {
	srv srvv1.Service
}

// NewItemAttributesController creates an item attributes handler.
func NewItemAttributesController(store store.Factory) *ItemAttributesController {
	return &ItemAttributesController{
		srv: srvv1.NewService(store),
	}
}

// Create creates a new item attribute.
func (iac *ItemAttributesController) Create(c *gin.Context) {
	var itemAttribute v1.ItemAttributes
	if err := c.ShouldBindJSON(&itemAttribute); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if err := iac.srv.ItemAttributes().Create(c, &itemAttribute, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, itemAttribute)
}

// Update updates an item attribute by its ID.
func (iac *ItemAttributesController) Update(c *gin.Context) {
	attributeID, _ := strconv.Atoi(c.Param("attributeID"))

	itemAttribute, err := iac.srv.ItemAttributes().Get(c, attributeID, metav1.GetOptions{})

	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	var newItemAttribute v1.ItemAttributes

	if err := c.ShouldBindJSON(&newItemAttribute); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	// Update itemAttribute fields here
	// itemAttribute.FieldName = newItemAttribute.FieldName

	if err := iac.srv.ItemAttributes().Update(c, itemAttribute, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, itemAttribute)
}

// Get retrieves an item attribute by its ID.
func (iac *ItemAttributesController) Get(c *gin.Context) {
	attributeID, _ := strconv.Atoi(c.Param("attributeID"))

	itemAttribute, err := iac.srv.ItemAttributes().Get(c, attributeID, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, itemAttribute)
}

// Delete deletes an item attribute by its ID.
func (iac *ItemAttributesController) Delete(c *gin.Context) {
	attributeID, _ := strconv.Atoi(c.Param("attributeID"))

	if err := iac.srv.ItemAttributes().Delete(c, attributeID, metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
