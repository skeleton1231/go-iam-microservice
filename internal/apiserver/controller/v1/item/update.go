package item

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
)

// Update updates an item by its ID.
func (ic *ItemController) Update(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("itemID"))

	var item v1.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = itemID

	if err := ic.srv.Items().Update(c, &item, metav1.UpdateOptions{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}
