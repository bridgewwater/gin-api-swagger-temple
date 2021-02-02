package router

import (
	"github.com/bridgewwater/gin-api-swagger-temple/handler/biz"
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

		// form
		// bizRouteGroup.GET("/form", biz.PostForm)
	}
}
