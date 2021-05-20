package params

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go-util/e"
	"go-util/util/logger"
	"go-util/util/response"
)

func InitAndValidParam(context *gin.Context, data interface{}) bool {
	err := context.ShouldBindBodyWith(data, binding.JSON)
	if err != nil {
		logger.L().Error("参数格式化异常：", err)
		response.Fail(context, *e.UserParamInvalid, nil)
		return false
	}
	// 参数校验
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		logger.L().Error("参数校验不通过：", err)
		response.Fail(context, *e.UserParamInvalid, nil)
		return false
	}
	return true
}
