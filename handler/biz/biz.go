package biz

import (
	"fmt"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary /biz/string
// @Description HealthCheck shows OK as the ping-pong result.
// @Tags biz
// @Success 200 "OK"
// @Router /biz/string [get]
func GetString(c *gin.Context) {
	message := "this is message"
	c.String(http.StatusOK, message)
}

type JsonBiz struct {
	Info   string `json:"info,omitempty" example:"input info here"`
	Id     string `json:"id,omitempty"  example:"id123zqqeeadg24qasd"`
	Offset int    `json:"offset,omitempty"  example:"0"`
	Limit  int    `json:"limit,omitempty"  example:"10"`
}

// @Summary /biz/json
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Success 200 {object} biz.JsonBiz "value in biz.JsonBiz"
// @Failure 403 {object} errdef.Err "error at errdef.Err"
// @Router /biz/json [get]
func GetJSON(c *gin.Context) {
	resp := JsonBiz{
		Info: "message",
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary /biz/path
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param some_id     path     string     true     "some id to show"
// @Success 200 {object} biz.JsonBiz "value in biz.JsonBiz"
// @Failure 403 {object} errdef.Err "error at errdef.Err"
// @Router /biz/path/{some_id} [get]
func GetPath(c *gin.Context) {
	id := c.Param("some_id")
	if id == "" {
		c.JSON(http.StatusForbidden, errdef.New(errdef.ErrParams, fmt.Errorf("id not found")))
		return
	}
	resp := JsonBiz{
		Id: id,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary /biz/query
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     false       "limit"
// @Success 200 {object} biz.JsonBiz "value in biz.JsonBiz"
// @Failure 403 {object} errdef.Err "error at errdef.Err"
// @Router /biz/query/ [get]
func GetQuery(c *gin.Context) {
	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrParams, err).Add("offset error"))
		return
	}
	limitStr := c.Query("limit")
	var limit int
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrParams, err).Add("limit error"))
			return
		}
	} else {
		limit = 10
	}
	resp := JsonBiz{
		Offset: offset,
		Limit:  limit,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary /biz/body
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param biz     body     biz.JsonBiz     true     "body biz.JsonBiz for post"
// @Success 200 {object} biz.JsonBiz "value in biz.JsonBiz"
// @Failure 400 {object} errdef.Err "error at errdef.Err"
// @Router /biz/body [post]
func PostBody(c *gin.Context) {
	var req JsonBiz
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrParams, err).Add("body error"))
		return
	}
	c.JSON(http.StatusOK, req)
}
