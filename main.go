package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	discoveryutil "github.com/bastien-mrq/service-a/discoveryUtil"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {

	response := discoveryutil.Register()

	log.Info("Service " + response.Uuid)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "service-a response",
			"uuid":    response.Port,
		})
		return
	})

	go func() {
		r.Run("127.0.0.1:" + strconv.Itoa(response.Port))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	discoveryutil.UnregisterService(response.Uuid)
}
