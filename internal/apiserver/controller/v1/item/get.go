package item

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// Get retrieves an item by its ID.
func (ic *ItemController) Get(c *gin.Context) {

	itemID, _ := strconv.Atoi(c.Param("itemID"))

	item, err := ic.srv.Items().Get(c, itemID, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}
