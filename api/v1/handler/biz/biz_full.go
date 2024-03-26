package biz

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/errdef"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/internal/parse_http"
	"github.com/gin-gonic/gin"
)

// GetHeaderFull
//
//	@Summary		/biz/header_full
//	@Description	for biz
//	@Tags			biz
//	@Accept			text/plain
//	@Produce		text/plain
//
//	@Param			BIZ_FOO		header		string				true	"header BIZ_FOO "	default(foo)
//	@Param			BIZ_BAR		header		string				true	"header BIZ_BAR "	default(bar)
//
//	@Success		200			{object}	parse_http.HeaderContent			"value in parse_http.HeaderContent"
//	@Failure		400			{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/header_full								[get]
func GetHeaderFull(c *gin.Context) {
	header, err := parse_http.HttpHeader(c)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err)
		return
	}
	handler.JsonSuccess(c, header)
}

// GetQueryFull
//
//	@Summary		/biz/query_full
//	@Description	for biz
//	@Tags			biz
//	@Accept			json
//	@Produce		json
//
//	@Param			foo		query		string				true	"params foo "	default("")
//	@Param			bar		query		string				true	"params bar "	default("")
//	@Param			baz		query		string				true	"params baz "	default("")
//
//	@Success		200			{object}	parse_http.QueryContent			"value in parse_http.QueryContent"
//	@Failure		400			{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/query_full								[get]
func GetQueryFull(c *gin.Context) {
	query, err := parse_http.HttpQuery(c)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err)
		return
	}
	handler.JsonSuccess(c, query)
}

// PostFormFull
//
//	@Summary		/biz/form_full
//	@Description	for biz
//	@Tags			biz
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//
//	@Param			foo		formData		string				true	"form item foo"		default(foo)
//	@Param			bar		formData		string				true	"form item bar"		default(bar)
//	@Param			baz		formData		string				true	"form item baz"		default(baz)
//
//	@Success		200			{object}	parse_http.FormContent			"value in parse_http.FormContent"
//	@Failure		400			{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/form_full								[post]
func PostFormFull(c *gin.Context) {
	query, err := parse_http.FormPost(c)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err)
		return
	}
	handler.JsonSuccess(c, query)
}
