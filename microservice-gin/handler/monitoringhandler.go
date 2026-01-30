package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
define and initialise monitoring APIs endpoints
*/
func InitializeMonitoringRoutes(publicGroupRouter *gin.RouterGroup) {
	log.Println("monitoring initialization..")
	publicGroupRouter.GET("/health", health)
	log.Println("monitoring initialized")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "Up/Running",
		"timestamp": time.Now().String(),
	})
}
