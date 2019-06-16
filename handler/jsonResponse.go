package handler

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/model"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
	"net/http"
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
