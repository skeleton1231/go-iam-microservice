package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var httpServer *http.Server
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiCmd",
	Long:  `apiCmd support apiserver`,
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%v:%v", "0.0.0.0", 9888)
		engine := gin.New()

		engine.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})

		httpServer = &http.Server{
			Addr:        address,
			Handler:     engine,
			IdleTimeout: time.Minute,
		}
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server start Fail", err)
		}
	},
}
