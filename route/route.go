package route

import (
	"ticketsale/controller"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	route.POST("/login", controller.Login)
	route.POST("/register", controller.Register)
	api := route.Use()
	// ticket
	api.POST("/ticket/create", controller.CreateTicket)
	api.POST("/ticket/update/:id", controller.UpdateTicket)
	api.DELETE("/ticket/delete/:id", controller.DeleteTicket)
	// ticket type
	api.POST("/ticket_type/create", controller.CreateTicketType)
	api.POST("/ticket_type/update/:id", controller.UpdateTicketType)
	api.DELETE("/ticket_type/delete/:id", controller.DeleteTicketType)
	// ticket sale price
	api.POST("/ticket_sale_price/create", controller.CreateTicketSalePrice)
	api.POST("/ticket_sale_price/update/:id", controller.UpdateTicketSalePrice)
	api.DELETE("/ticket_sale_price/delete/:id", controller.DeleteTicketSalePrice)
}
