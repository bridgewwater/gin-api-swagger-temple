package biz

import (
	"net/http"

	"github.com/bridgewwater/gin-api-swagger-temple/handler"
	"github.com/bridgewwater/gin-api-swagger-temple/model/biz"
	"github.com/bridgewwater/gin-api-swagger-temple/pkg/errdef"
	"github.com/gin-gonic/gin"
)

// @Summary /biz/json
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Success    200    {object}    biz.Biz    "value in biz.Biz"
// @Failure    500
// @Router /biz/json [get]
func GetJSON(c *gin.Context) {
	resp := biz.Biz{
		Info: "message",
	}
	handler.JsonSuccess(c, struct {
		NewInfo string `json:"new_info"`
	}{NewInfo: resp.Info})
}

// @Summary /biz/modelBiz
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept application/json
// @Produce application/json
// @Param    biz    body    biz.Biz    true    "body biz.Biz for post"
// @Success    200    {object}    biz.Biz    "value in biz.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/modelBiz [post]
func PostJsonModelBiz(c *gin.Context) {
	var req biz.Biz
	if err := c.BindJSON(&req); err != nil {
		handler.JsonErrDefErr(c, errdef.ErrBind, err, "limit error")
		c.JSON(http.StatusBadRequest, errdef.New(errdef.ErrBind, err).Add("body error"))
		return
	}
	c.JSON(http.StatusOK, req)
}
