// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetItemController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFactory := new(store.MockFactory)
	mockItemStore := new(store.MockItemStore)
	mockFactory.On("Items").Return(mockItemStore) // Set the return value for Items method
	mockFactory.On("Close").Return(nil)

	mockService := new(v1.MockService)
	mockItemService := new(v1.MockItemService)
	mockService.On("Item").Return(mockItemService)

	itemID := 1000000000
	mockItem := &model.Item{
		ASIN:  "B09XN3WCVG",
		Title: "Item 1",
	}

	mockItemStore.On("Get", mock.Anything, itemID, metav1.GetOptions{}).Return(mockItem, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+strconv.Itoa(itemID), nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/api/v1/items/:itemID", NewItemController(mockFactory).Get)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
