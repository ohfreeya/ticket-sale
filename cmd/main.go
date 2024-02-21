package main

import (
	"ticketsale/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main(){
	server := gin.Default()
}