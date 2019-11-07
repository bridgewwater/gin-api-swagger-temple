package biz

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/handler"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
)

func PostForm(c *gin.Context) {
	c.GetHeader("")
	r := c.Request
	err := r.ParseForm()
	if err != nil {
		handler.JsonErrDef(c, errdef.ErrParse, "Form parse error")
		return
	}
}
