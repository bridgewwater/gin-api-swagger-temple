package biz

import "github.com/gin-gonic/gin"

func Router(g *gin.Engine, basePath string) {
	bizRouteGroup := g.Group(basePath + "/biz")
	{
		bizRouteGroup.GET("/string", GetString)
		bizRouteGroup.GET("/path/:some_id", GetPath)
		bizRouteGroup.GET("/query", GetQuery)
		bizRouteGroup.GET("/json", GetJSON)

		// post

		bizRouteGroup.POST("/form", PostForm)
		bizRouteGroup.POST("/modelBiz", PostJsonModelBiz)
		bizRouteGroup.POST("/modelBizQuery", PostQueryJsonMode)
	}
}
