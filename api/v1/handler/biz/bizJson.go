package biz

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/errdef"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/model/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetJSON
//
//	@Summary		api demo json
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			application/json
//	@Produce		application/json
//
//	@Success		200		{object}		biz.Biz			"value in biz.Biz"
//	@Failure		500										""
//	@Router			/biz/json								[get]
func GetJSON(c *gin.Context) {
	resp := biz.Biz{
		Info: "message",
	}
	handler.JsonSuccess(c, struct {
		NewInfo string `json:"new_info"`
	}{NewInfo: resp.Info})
}

// PostJsonModelBiz
//
//	@Summary		api json struct biz.Biz
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			application/json
//	@Produce		application/json
//
//	@Param			biz		body			biz.Biz			true	"body biz.Biz for post"
//	@Success		200		{object}		biz.Biz			"value in biz.Biz"
//
//	@Failure		400		{object}		errdef.Err		"error at errdef.Err"
//	@Router			/biz/modelBiz							[post]
func PostJsonModelBiz(c *gin.Context) {
	var req biz.Biz
	if err := c.BindJSON(&req); err != nil {
		handler.JsonErrDefErr(c, errdef.ErrBind, err, "limit error")
		c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrBind, err).Add("body error"))
		return
	}
	c.JSON(http.StatusOK, req)
}

// PostQueryJsonMode
//
//	@Summary		api post query with json struct biz.Biz
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			application/json
//	@Produce		application/json
//
//	@Param			offset	query			int			true	"Offset"
//	@Param			limit	query			int			false	"limit"
//	@Param			biz		body			biz.Biz		true	"body biz.Biz for post"
//
//	@Success		200		{object}		biz.Biz		"value in biz.Biz"
//	@Failure		400		{object}		errdef.Err	"error at errdef.Err"
//	@Router			/biz/modelBizQuery					[post]
func PostQueryJsonMode(c *gin.Context) {
	offset, limit, err := handler.ParseQueryCommonOffsetAndLimit(c)
	if err != nil {
		handler.JsonErrDefErr(c, errdef.ErrParams, err)
		return
	}
	var req biz.Biz
	if errBind := c.BindJSON(&req); errBind != nil {
		handler.JsonErrDefErr(c, errdef.ErrBind, errBind)
		return
	}
	req.Offset = offset
	req.Limit = limit
	c.JSON(http.StatusOK, req)
}
