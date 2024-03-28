package tests

import (
	"gin-cmdb/config"
	"gin-cmdb/data"
	"gin-cmdb/internal/pkg/logger"
	"gin-cmdb/internal/routers"
	"gin-cmdb/internal/validator"
	"gin-cmdb/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
	"strings"
)

func SetupRouter() *gin.Engine {
	// 1、初始化配置
	config.InitConfig("")
	config.Config.Mysql.PrintSql = false
	// 2、初始化zap日志
	logger.InitLogger()
	// 初始化数据库
	data.InitData()
	// 初始化验证器
	validator.InitValidatorTrans("zh")

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	engine := gin.Default()

	routers.SetAdminApiRoute(engine)
	return engine
}

func Request(method, route string, body *string, args ...any) *utils.HttpRequest {
	h := utils.HttpRequest{}
	var params io.Reader
	if body != nil {
		params = strings.NewReader(*body)
	}

	return h.JsonRequest(method, route, params, args...)
}

func GetRequest(route string, queryParams *url.Values, args ...any) *utils.HttpRequest {
	h := utils.HttpRequest{}

	return h.GetRequest(route, queryParams, args...)
}
