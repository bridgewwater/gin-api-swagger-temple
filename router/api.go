package router

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/handler/biz"
	"github.com/gin-gonic/gin"
)

func bizApi(g *gin.Engine, basePath string) {
	// The health check handlers
	bizRouteGroup := g.Group(basePath + "/biz")
	{
		bizRouteGroup.POST("/modelBiz", biz.PostJsonModelBiz)
		bizRouteGroup.GET("/string", biz.GetString)
		bizRouteGroup.GET("/json", biz.GetJSON)
		bizRouteGroup.GET("/path/:some_id", biz.GetPath)
		bizRouteGroup.GET("/query", biz.GetQuery)
	}
}
