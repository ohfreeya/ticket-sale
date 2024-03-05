package middleware

import "github.com/gin-gonic/gin"

func JWTAuth(ctx *gin.Context) {
	// auth the token 
	// if not pass, return 401
	

	// if pass, continue
	ctx.Next()
}