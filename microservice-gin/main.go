package main

import (
	"log"
	"microservice-gin/server"
)

func main() {
	log.Println("app starting..")
	server.StartServer()
}
