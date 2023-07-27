package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Description: 响应结构体
 * @param ctx gin.Context
 * @param httpStatus http状态码
 * @param code 自定义状态码
 * @param data 数据
 * @param msg 提示信息
 */
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

/**
 * @Description: 成功响应
 * @param ctx gin.Context
 * @param data 数据
 * @param msg 提示信息
 */
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

/**
 * @Description: 失败响应
 * @param ctx gin.Context
 * @param msg 提示信息
 * @param data 数据
 */
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
