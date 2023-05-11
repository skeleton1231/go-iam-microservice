package item

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListItemController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFactory := new(store.MockFactory)
	mockItemStore := new(store.MockItemStore)
	mockFactory.On("Items").Return(mockItemStore)
	mockFactory.On("Close").Return(nil)

	itemList := &model.ItemList{
		Items: []*model.Item{
			{
				ID:    1000000000,
				ASIN:  "B09XN3WCVG",
				Title: "Item 1",
			},
			{
				ID:    1000000001,
				ASIN:  "B09XN3WCVH",
				Title: "Item 2",
			},
		},
	}

	mockItemStore.On("List", mock.Anything, mock.AnythingOfType("v1.ListOptions")).Return(itemList, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/api/v1/items", NewItemController(mockFactory).List)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseItems model.ItemList
	err := json.Unmarshal(w.Body.Bytes(), &responseItems)
	assert.NoError(t, err)

	assert.Equal(t, len(itemList.Items), len(responseItems.Items))
	for i, item := range itemList.Items {
		assert.Equal(t, item.ID, responseItems.Items[i].ID)
		assert.Equal(t, item.ASIN, responseItems.Items[i].ASIN)
		assert.Equal(t, item.Title, responseItems.Items[i].Title)
	}
}
