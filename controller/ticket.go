package controller

import (
	"net/http"
	"ticketsale/model"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTicket(ctx *gin.Context) {
	type form struct {
		Name      string `json:"name"`
		Introduce string `json:"introduce"`
		Count     int    `json:"count"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	var model model.Tickets
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Name = req.Name
	model.Introduce = req.Introduce
	model.Count = req.Count
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	model.ExpiresAt = expireAt
	err = model.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Create ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create ticket success",
	})
}

func UpdateTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	type form struct {
		Count int `json:"count"`
		Introduce string `json:"introduce"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	var model model.Tickets
	_, err :=  model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Count = req.Count
	model.Introduce = req.Introduce
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	model.ExpiresAt = expireAt
	err = model.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Update ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update ticket success",
	})

}

func DeleteTicket(ctx *gin.Context) {

}
