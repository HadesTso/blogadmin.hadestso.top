package common

import (
	"github.com/gin-gonic/gin"
	"newbee.hadestso.top/routes"
)

func SetupRouter(router *gin.Engine) {

	// 注册后台 Admin 路由
	routes.RegisterAdminRouter(router)

	// 注册 API 路由
	routes.RegisterAPIRouter(router)

}
