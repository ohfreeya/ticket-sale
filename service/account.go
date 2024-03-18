package service

import (
	"strings"
	"ticketsale/auth"

	"github.com/gin-gonic/gin"
)

func GetUidFromToken(ctx *gin.Context) int {
	token := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
	userInfo, _ := auth.VerifyToken(token[1])
	return userInfo.Uid

}
