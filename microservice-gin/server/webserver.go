package server

import (
	"log"
	"microservice-gin/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("building gin router..")
	router := gin.Default()
	log.Println("gin router created", router)
	handler.InitializeRoutes(router)
	router.Run("0.0.0.0:8090")
}
