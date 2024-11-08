package req

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个泛型的统一响应结构体
type Req[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// 定义一个工厂函数来创建成功的响应
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Req[T]{
		Code:    200,
		Message: "Success!",
		Data:    data,
	})
}

// 定义一个工厂函数来创建错误的响应
func Err[T any](c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Req[T]{
		Code:    code,
		Message: message,
		Data:    *new(T), // 使用零值初始化 Data 字段
	})
}

