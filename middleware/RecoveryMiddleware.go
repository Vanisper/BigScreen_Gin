package middleware

import (
	"fmt"
	"log"

	"BigScreen_Gin/response"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获panic异常
func RecoveryMiddleware() gin.HandlerFunc {
	log.Println("捕获panic异常中间件挂载成功")
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()

		ctx.Next()
	}
}
