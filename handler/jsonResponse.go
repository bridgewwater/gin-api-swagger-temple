package handler

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/model"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// use as
//	handler.JsonSuccess(c)
//	return
func JsonSuccess(c *gin.Context, data ...interface{}) {
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

// use as
//	handler.JsonErrDef(c, http.StatusBadRequest, errdef.ErrParams)
//	return
// or use add messages, sep of message use string "; "
//	handler.JsonErrDef(c, http.StatusBadRequest, errdef.ErrParams, "id", "not found, set id and retry")
//	return
func JsonErrDef(c *gin.Context, httpCode int, def *errdef.ErrDef, errMsgs ...string) {
	if httpCode == 0 {
		httpCode = http.StatusGone
	}
	err := errdef.NewErr(def)
	if len(errMsgs) == 0 {
		c.JSON(httpCode, err)
		return
	} else {
		message := strings.Join(errMsgs, "; ")
		c.JSON(httpCode, err.Add(message))
	}

}

// use as
//	handler.JsonErr(c, 0)
//	return
func JsonErr(c *gin.Context, errCode int, data ...interface{}) {
	if errCode == 0 {
		errCode = errdef.InternalServerError.Code
	}
	if data != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: errCode,
			Msg:  "success",
			Data: data,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Code: errCode,
			Msg:  "success",
		})
	}
}
