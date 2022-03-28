package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRouter(r *gin.Engine)  {
	admin := r.Group("/v1")

	// 注册一个路由
	admin.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
}