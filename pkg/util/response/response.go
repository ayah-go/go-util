package response

import (
	"go-util/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	//Context.Header("key2021","value2021")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
// func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
// 	Context.Header("Content-Type", "application/json; charset=utf-8")
// 	Context.String(httpCode, jsonStr)
// }

// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, e.SUCCESS.Code, e.SUCCESS.Msg, data)
}

// 失败的业务逻辑
func Fail(c *gin.Context, enum e.ErrorCode, data interface{}) {
	Response(c, http.StatusBadRequest, enum.Code, enum.Msg, data)
	c.Abort()
}
