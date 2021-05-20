package aop

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-util/util/json"
	"go-util/util/logger"
	"strings"
	"time"
)

func LoggerHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now() // Next 在这里相当于 接口函数，在Next之前则在接口函数之前执行

		random, err := uuid.NewRandom()
		if err != nil {
			logger.L().Error("生成UUID error : ", err)
		}

		var body string

		//	Get请求不需要打印 body
		if strings.EqualFold("GET", strings.ToUpper(context.Request.Method)) {
			logger.L().Infof("  %s  %s  【请求调入】 method : %s  IP : %s  FORM数据 : %s", context.Request.Method, random,
				context.Request.RequestURI, context.Request.RemoteAddr, json.StructToJSON(context.Request.Form))
		} else {
			err := context.BindJSON(body)
			if err != nil {
				body = ""
			}

			logger.L().Infof("  %s  %s  【请求调入】 method : %s  IP : %s  请求数据 : %s", context.Request.Method, random,
				context.Request.RequestURI, context.Request.RemoteAddr, body)
		}

		context.Next()

		cost := time.Since(start) // Next 之后，则相当于在接口函数之后执行，形成了一个切面
		logger.L().Infof(" %s  【返回结果】 method : %s, 请求用时 : %dms , 返回结果 : %s", random,
			context.Request.RequestURI, cost.Milliseconds(), json.StructToJSON(context.Request.Response))

		return
	}
}
