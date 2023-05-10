package item

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockItemService) Create(ctx *gin.Context, item *model.Item) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func TestCreateItemController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(v1.MockService)
	mockItemService := new(MockItemService)
	mockService.On("Item").Return(mockItemService)

	item := &model.Item{
		ASIN: "B123456",
		// Add other fields here
	}

	mockItemService.On("Create", mock.Anything, item).Return(nil)

	data, _ := json.Marshal(item)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	r := gin.Default()
	mockFactory := new(store.MockFactory)
	mockItemStore := new(store.MockItemStore)
	mockFactory.On("Item").Return(mockItemStore)

	r.POST("/api/v1/items", NewItemController(mockFactory).Create)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
