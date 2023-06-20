package biz

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/errdef"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/model/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPath
//
//	@Summary		api path demo for route path
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			some_id		path		string			true	"some id to show"
//
//	@Success		200			{object}	biz.Biz			"value in biz.Biz"
//	@Failure		400			{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/path/{some_id}						[get]
func GetPath(c *gin.Context) {
	id := c.Param("some_id")
	if id == "" {
		handler.JsonErrDef(c, errdef.ErrParams, "id not found")
		return
	}
	resp := biz.Biz{
		Id: id,
	}
	handler.JsonSuccess(c, resp)
}

// GetQuery
//
//	@Summary		api demo for query.
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			json
//	@Produce		json
//
//	@Param			offset		query		int				true	"Offset"
//	@Param			limit		query		int				false	"limit"
//
//	@Success		200			{object}	biz.Biz			"value in biz.Biz"
//	@Failure		400			{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/query/								[get]
func GetQuery(c *gin.Context) {
	offset, limit, err := handler.ParseQueryCommonOffsetAndLimit(c)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err)
		return
	}
	resp := biz.Biz{
		Offset: offset,
		Limit:  limit,
	}
	handler.JsonSuccess(c, resp)
}

// GetString
//
//	@Summary		sample demo string
//	@Description	get string of this api. warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			json
//	@Produce		plain
//
//	@Success		200										"OK"
//	@Failure		500										""
//	@Router			/biz/string								[get]
func GetString(c *gin.Context) {
	message := "this is biz message"
	c.String(http.StatusOK, message)
}
