package controller

import (
	"net/http"
	"ticketsale/model"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var req struct {
		UserID   uint    `json:"user_id"`
		TicketID uint    `json:"ticket_id"`
		Count    int     `json:"count"`
		Price    float64 `json:"price"`
		Total    float64 `json:"total"`
		Coupon   string  `json:"coupon"`
	}
	var model model.Orders
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.UserID = req.UserID
	model.TicketID = req.TicketID
	model.Count = req.Count
	model.Price = req.Price
	model.Total = req.Total
	model.Coupon = req.Coupon
	if err := model.Create(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Failed to create order",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		UserID   uint    `json:"user_id"`
		TicketID uint    `json:"ticket_id"`
		Count    int     `json:"count"`
		Price    float64 `json:"price"`
		Total    float64 `json:"total"`
		Coupon   string  `json:"coupon"`
	}
	var model model.Orders
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	target, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Order not found",
			"error": err.Error(),
		})
		return
	}
	target.UserID = req.UserID
	target.TicketID = req.TicketID
	target.Count = req.Count
	target.Price = req.Price
	target.Total = req.Total
	target.Coupon = req.Coupon
	if err := target.Update(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Failed to update order",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})

}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.Orders
	target, err := model.Find(map[string]interface{}{"id": id})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Order not found",
			"error": err.Error(),
		})
		return
	}

	if err := target.Delete(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Failed to delete order",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
