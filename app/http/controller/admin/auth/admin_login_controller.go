package auth

import (
	"github.com/gin-gonic/gin"
	"newbee.hadestso.top/app/http/controller/admin"
)

type LoginController struct {
	admin.BaseAdminController
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {
}

func (lc *LoginController) LoginByEmail(c *gin.Context) {
}
