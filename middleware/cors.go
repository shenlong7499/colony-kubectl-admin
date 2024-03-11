package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	method := ctx.Request.Method
	origin := ctx.Request.Header.Get("Origin")
	// 设置允许跨域请求的来源
	ctx.Header("Access-Control-Allow-Origin", origin)
	// 设置允许跨域请求的请求头
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization")
	// 设置允许跨域请求的请求方法
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	// 设置服务器支持的响应头字段
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// 设置是否允许发送 Cookie
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		// 如果是OPTIONS请求，表示允许客户端发送实际请求
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	// 继续处理请求
	ctx.Next()
}
