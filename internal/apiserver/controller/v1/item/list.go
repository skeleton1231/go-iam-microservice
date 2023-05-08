package item

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

// List handles the request to list all items.
func (ic *ItemController) List(c *gin.Context) {
	// Parse query parameters
	opts := metav1.ListOptions{}

	// if filter := c.Query("filter"); filter != "" {
	// 	opts.Filter = filter
	// }

	// Call the List method from the ItemSrv interface
	items, err := ic.srv.Items().List(c.Request.Context(), opts)
	if err != nil {
		// Handle error
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	// Write the response with the items list
	core.WriteResponse(c, nil, items)
}
