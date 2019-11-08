package biz

import (
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/handler"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/pkg/errdef"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Summary /biz/form
// @Description warning api in prod will hide, abs remote api for dev
// @Tags biz
// @Accept  application/x-www-form-urlencoded
// @Produce application/x-www-form-urlencoded
// @Param    biz    form    model.Biz    true    "body model.Biz for post"
// @Success    200    {object}    model.Biz    "value in model.Biz"
// @Failure    400    {object}    errdef.Err    "error at errdef.Err"
// @Router /biz/form [post]
func PostForm(c *gin.Context) {
	c.GetHeader("")
	r := c.Request
	err := r.ParseForm()
	if err != nil {
		handler.JsonErrDef(c, errdef.ErrParse, "Form parse error")
		return
	}
	formContent := make(map[string]string)
	for k, v := range r.PostForm {
		formContent[k] = strings.Join(v, "")
	}
	handler.JsonSuccess(c, struct {
		PostFormContent map[string]string `json:"post_form_content,omitempty"`
	}{
		PostFormContent: formContent,
	})
}
