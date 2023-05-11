package item

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteItemController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFactory := new(store.MockFactory)
	mockItemStore := new(store.MockItemStore)
	mockFactory.On("Items").Return(mockItemStore)
	mockFactory.On("Close").Return(nil)

	itemID := 1000000000

	mockItemStore.On("Delete", mock.Anything, itemID, metav1.DeleteOptions{}).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+strconv.Itoa(itemID), nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/api/v1/items/:itemID", NewItemController(mockFactory).Delete)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
