package middleware

import (
	"strings"
	"ticketsale/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuth(ctx *gin.Context) {
	// auth the token
	token := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
	if len(token) != 2 {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "Invalid token",
		})
		ctx.Abort()
		return
	}
	// verify the token
	_, err := auth.VerifyToken(token[1])
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "Invalid token",
		})
		ctx.Abort()
		return
	}
	// if pass, continue
	ctx.Next()
}
