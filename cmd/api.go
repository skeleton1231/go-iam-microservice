package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var httpServer *http.Server
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiCmd",
	Long:  `apiCmd support apiserver`,
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%v:%v", "0.0.0.0", 9888)
		gin.SetMode(gin.ReleaseMode)
		h2s := &http2.Server{}
		engine := gin.New()

		engine.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})

		httpServer = &http.Server{
			Addr:        address,
			Handler:     h2c.NewHandler(engine, h2s),
			IdleTimeout: time.Minute,
		}

		quitChan := make(chan os.Signal)
		signal.Notify(quitChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		go func() {
			err := httpServer.ListenAndServe()
			if err != nil {
				logrus.Error(err.Error())
				os.Exit(0)
			}
		}()

		//退出应用
		<-quitChan
		logrus.Info("Server exit")
		_ = httpServer.Close()

		// if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	fmt.Println("Server start Fail", err)
		// }
	},
}
