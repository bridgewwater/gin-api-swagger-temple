package biz

import (
	"net/http"
	"strconv"

	"github.com/bridgewwater/gin-api-swagger-temple/handler"
	"github.com/bridgewwater/gin-api-swagger-temple/model/biz"
	"github.com/bridgewwater/gin-api-swagger-temple/pkg/errdef"
	"github.com/gin-gonic/gin"
)

// @Summary /biz/path
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param    some_id    path    string    true    "some id to show"
// @Success    200    {object}    biz.Biz    "value in biz.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/path/{some_id} [get]
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

// @Summary /biz/query
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param    offset    query    int    true    "Offset"
// @Param    limit    query    int    false    "limit"
// @Success    200    {object}    biz.Biz    "value in biz.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/query/ [get]
func GetQuery(c *gin.Context) {
	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err, "offset error")
		return
	}
	limitStr := c.Query("limit")
	var limit int
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			handler.JsonErrDefErr(c, errdef.ErrParams, err, "limit error")
			return
		}
	} else {
		limit = 10
	}
	resp := biz.Biz{
		Offset: offset,
		Limit:  limit,
	}
	handler.JsonSuccess(c, resp)
}

// @Summary /biz/string
// @Description get string of this api.
// @Tags biz
// @Success    200    "OK"
// @Failure    500
// @Router /biz/string [get]
func GetString(c *gin.Context) {
	message := "this is biz message"
	c.String(http.StatusOK, message)
}
