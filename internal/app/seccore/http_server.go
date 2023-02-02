package seccore

import (
	"github.com/device-security-v1/seccore/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func httpServer() server.Server {
	// 生成一个 gin.Engine 实例，即：WSGI 应用程序
	e := gin.New()

	// 添加会使用到的中间件
	e.Use(gin.Recovery())

	// 添加路由
	route(e)

	// 返回封装的统一的 Server 接口实例，这个 Server实例最终会被放在 ServerGroup中来统一启停
	return server.WrapGin(e, ":"+viper.GetString("config.port"))
}

func route(e *gin.Engine) {
	// URL: /v1/sec/report
	v1rg := e.Group("/v1/sec")
	v1rg.POST("report", _reportService.HandleReportV1)
}
