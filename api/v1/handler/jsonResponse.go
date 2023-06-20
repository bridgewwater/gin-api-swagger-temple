package handler

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/errdef"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JsonErrDef
// use as
// handler.JsonErrDef(c, errdef.ErrParams)
// return
// or use add messages, sep of message use string "; "
// handler.JsonErrDef(c, errdef.ErrParams, "id", "not found, set id and retry")
// return
func JsonErrDef(c *gin.Context, def *errdef.ErrDef, errMsgs ...string) {
	err := errdef.NewErr(def)
	if len(errMsgs) == 0 {
		c.JSON(def.HttpStatus, err)
		return
	} else {
		message := strings.Join(errMsgs, "; ")
		c.JSON(def.HttpStatus, err.Add(message))
	}

}

// JsonErrDefErr
// use as
// handler.JsonErrDefErr(c, errdef.ErrDatabase, err)
// return
// or use add messages, sep of message use string "; "
// handler.JsonErrDefErr(c, errdef.ErrDatabase, err, "can not found")
// return
func JsonErrDefErr(c *gin.Context, def *errdef.ErrDef, err error, errMsg ...string) {
	errResp := errdef.New(def, err)
	if len(errMsg) == 0 {
		c.JSON(def.HttpStatus, errResp)
		return
	} else {
		message := strings.Join(errMsg, "; ")
		c.JSON(def.HttpStatus, errResp.Add(message))
	}

}

// JsonSuccess
// use as
// handler.JsonSuccess(c)
// return
func JsonSuccess(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Code: 0,
			Msg:  "success",
		})
	}
}

// JsonFail
// use as
// handler.JsonFail(c, "fail message")
// return
func JsonFail(c *gin.Context, fail string, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: 0,
			Msg:  fail,
			Data: data,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Code: 0,
			Msg:  fail,
		})
	}
}

// JsonErr
// use as
// handler.JsonErr(c, 0)
// return
func JsonErr(c *gin.Context, errCode int, data interface{}) {
	if errCode == 0 {
		errCode = errdef.InternalServerError.Code
	}
	if data != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: errCode,
			Msg:  "error",
			Data: data,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Code: errCode,
			Msg:  "error",
		})
	}
}
