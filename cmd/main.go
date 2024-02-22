package main

import (
	"ticketsale/config"
	"ticketsale/model"
	"ticketsale/route"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	server := gin.Default()
	err := config.DBInit()
	if err != nil {
		panic(err)
	}
	model.Migrate()
	route.Route(server)
	server.Run(":8080")
}
