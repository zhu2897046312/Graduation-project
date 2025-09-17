package utils

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 状态码，0表示成功
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"result"`    // 返回数据
	Error   interface{} `json:"error"`    
	Time    int64       `json:"time"`    
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
		Error: nil,
	})
}

// 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// 参数错误响应
func InvalidParams(c *gin.Context) {
	Error(c, 1001, "参数错误")
}

func InvalidParams_1(c *gin.Context,req interface{}) {
	fmt.Println(req)
	Error(c, 1001, "参数错误")
}

// 服务器错误响应
func ServerError(c *gin.Context, err error) {
	// 实际项目中可以记录日志
	Error(c, 500, "服务器内部错误: "+err.Error())
}