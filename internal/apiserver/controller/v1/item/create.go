package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
)

// Create creates a new item.
func (ic *ItemController) Create(c *gin.Context) {
	var item v1.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ic.srv.Items().Create(c, &item, metav1.CreateOptions{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}
