package main

import (
	"github.com/gin-gonic/gin"
	"res-api.com/apis/db"
	"res-api.com/apis/routes"
)

func main() {
	db.InitDb()
	server := gin.Default() // create a new server
	routes.RegisterRoutes(server)

	server.Run(":8080") // http://localhost:8080

}
