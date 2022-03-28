package response

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "操作成功！",
	})
}

// Data 响应 200 和带 data 键的 JSON 数据
func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data": data,
	})
}

// Created 响应 201 和带 data 键的 JSON 数据
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": data,
	})
}

// CreatedJSON 响应 201 和带 data 键的 JSON 数据
func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("权限不足，请确认您有对应的权限", msg...),
	})
}

// Abort404 响应 404，数据无法查找时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("数据不存在，请确认请求正确", msg...),
	})
}

// Abort500 响应 500，服务器内部错误时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}

// BadRequest 响应 400, 错误的请求方式时使用默认消息
func BadRequest(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// Error 响应 404 或 422
func Error(c *gin.Context, err error, msg ...string) {
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("请求处理失败，请查看 error 的值", msg...),
		"error": err.Error(),
	})
}

// ValidationError 处理表单验证不通过的错误
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "请求验证不通过，具体请查看 errors",
		"errors": errors,
	})
}

// Unauthorized 响应 401，未传参 msg 时使用默认消息
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// defaultMessage 内部使用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}

	return message
}