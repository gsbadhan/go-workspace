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
	engine := gin.Default()
	log.Println("gin router created", engine)
	publicGroupRouter := engine.Group("/public")
	privateGroupRouter := engine.Group("/private")

	// enable or disable JWT security
	handler.EnableJwtSecurity(engine, publicGroupRouter, privateGroupRouter)
	// enable or disable album routes
	handler.InitializeAlbumRoutes(privateGroupRouter)
	// enable or disable health check routes
	handler.InitializeMonitoringRoutes(privateGroupRouter)

	err := engine.Run("0.0.0.0:8090")
	panic(err)
}
