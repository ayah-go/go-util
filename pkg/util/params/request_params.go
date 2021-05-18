package params

import (
	"go-util/pkg/e"
	"go-util/pkg/util/logger"
	"go-util/pkg/util/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
