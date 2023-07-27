package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	log.Println("跨域中间件挂载成功")
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
