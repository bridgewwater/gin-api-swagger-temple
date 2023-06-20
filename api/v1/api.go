package v1

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler/biz"
	"github.com/gin-gonic/gin"
)

func bizApi(g *gin.Engine, basePath string) {

	biz.Router(g, basePath)
	// TODO sinlov 2023/6/20 other router group at here
}
