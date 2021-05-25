package check_token

import (
	"github.com/ayah-go/go-util/e"
	"github.com/ayah-go/go-util/util/ast"
	"github.com/ayah-go/go-util/util/logger"
	"github.com/ayah-go/go-util/util/response"
	"github.com/gin-gonic/gin"
)

var (
	urlList       []string
	targetComment = "@CheckToken"
)

/*
CheckToken 路由中间件
@author wyy
@date 2021/4/22 13:59
*/
func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.L().Info(urlList)
		needCheckToken := false
		for _, url := range urlList {
			if url == context.FullPath() {
				needCheckToken = true
				break
			}
		}
		if needCheckToken {
			// 校验token
			logger.L().Info("需要鉴权")
			token := context.GetHeader("access-token")
			if token == "" {
				response.Fail(context, *e.UserSysUserNotExists, nil)
				context.Abort()
			}
			//todo redis 鉴权
		}
		context.Next()
		return

	}
}

/*
InitCheckToken 扫描带有@CheckToken注释的路由并提取url参数，用于token鉴权
@author wyy
@date 2021/4/22 13:59
*/
func InitCheckToken(apiPath string) {
	// 遍历api文件夹下的所有go文件
	urlList = ast.GetUrlListFromControllerPathByComment(apiPath, targetComment)
	logger.L().Info("以下URL需要TOKEN鉴权:", urlList)
}
