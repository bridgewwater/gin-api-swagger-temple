package biz

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/handler"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/model"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary /biz/one
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param    biz    body    model.Biz    true    "body model.Biz for post"
// @Success    200    {object}    model.Biz    "value in model.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/one [post]
func PostOne(c *gin.Context) {
	var req model.Biz
	if err := c.BindJSON(&req); err != nil {
		handler.JsonErrDefErr(c, errdef.ErrBind, err, "limit error")
		c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrBind, err).Add("body error"))
		return
	}
	c.JSON(http.StatusOK, req)
}

// @Summary /biz/path
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param    some_id    path    string    true    "some id to show"
// @Success    200    {object}    model.Biz    "value in model.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/path/{some_id} [get]
func GetPath(c *gin.Context) {
	id := c.Param("some_id")
	if id == "" {
		handler.JsonErrDef(c, errdef.ErrParams, "id not found")
		return
	}
	resp := model.Biz{
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
// @Success    200    {object}    model.Biz    "value in model.Biz"
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
	resp := model.Biz{
		Offset: offset,
		Limit:  limit,
	}
	handler.JsonSuccess(c, resp)
}

// @Summary /biz/json
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Success    200    {object}    model.Biz    "value in model.Biz"
// @Failure    500
// @Router /biz/json [get]
func GetJSON(c *gin.Context) {
	resp := model.Biz{
		Info: "message",
	}
	handler.JsonSuccess(c, struct {
		NewInfo string `json:"new_info"`
	}{NewInfo: resp.Info})
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
