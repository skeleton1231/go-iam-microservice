// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateItemController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFactory := new(store.MockFactory)
	mockItemStore := new(store.MockItemStore)
	mockFactory.On("Items").Return(mockItemStore) // Set the return value for Items method
	mockFactory.On("Close").Return(nil)

	mockService := new(v1.MockService)
	mockItemService := new(v1.MockItemService)
	mockService.On("Item").Return(mockItemService)

	item := &model.Item{
		ASIN:  "B09XN3WCVG",
		Title: "Bone Conduction Headphones, Bluetooth Wireless Running Headphone, Open Ear Earphone with Mic, Sweat Resistant Sport Headset for Running, Gym, Cycling, Walking, Workout, Hiking, Listening (Black)",
		// Add other fields here
	}

	mockItemStore.On("Update", mock.Anything, mock.AnythingOfType("*model.Item"), mock.AnythingOfType("v1.UpdateOptions")).Return(nil)

	data, _ := json.Marshal(item)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/items/"+strconv.Itoa(int(item.ID)), bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	r := gin.Default()
	r.PUT("/api/v1/items/:itemID", NewItemController(mockFactory).Update)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
