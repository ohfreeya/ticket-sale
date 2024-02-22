package route

import (
	"ticketsale/controller"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	route.POST("/login", func(ctx *gin.Context) {})
	route.POST("/register", controller.Register)
}
