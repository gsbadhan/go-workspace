package server

import (
	"log"
	"microservice-gin/handler"

	"github.com/gin-gonic/gin"
)

/*
Start HTTP web server
*/

func StartServer() {
	log.Println("building gin router..")
	router := gin.Default()
	log.Println("gin router created", router)
	handler.InitializeRoutes(router)
	handler.InitializeMonitoring(router)
	err := router.Run("0.0.0.0:8090")
	panic(err)
}
