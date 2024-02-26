package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTicket(ctx *gin.Context) {
	type form struct {
		Name string `json:"name"`
		Introduce string `json:"introduce"`
		Count int `json:"count"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid form",
			"error": err.Error(),
		})
		return
	}
}

func UpdateTicket(ctx *gin.Context) {

}

func DeleteTicket(ctx *gin.Context) {

}