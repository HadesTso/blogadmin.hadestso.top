package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"newbee.hadestso.top/common"
)

func main() {
	// gin 实例
	router := gin.New()

	// 初始化路由绑定
	common.SetupRouter(router)

	// 运行服务器
	err := router.Run(":5056")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
