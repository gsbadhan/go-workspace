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
	handler.InitializeAlbumRoutes(privateGroupRouter)
	handler.InitializeMonitoringRoutes(publicGroupRouter)
	err := engine.Run("0.0.0.0:8090")
	panic(err)
}
