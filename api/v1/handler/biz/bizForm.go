package biz

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/errdef"
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/handler"
	"github.com/gin-gonic/gin"
	"strings"
)

// PostForm
//
//	@Summary		api demo form with: x-www-form-urlencoded
//	@Description	warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//
//	@Param			biz		formData	biz.Biz			true	"body model.Biz for post"
//
//	@Success		200		{object}	biz.Biz			"value in model.Biz"
//	@Failure		400		{object}	errdef.Err		"error at errdef.Err"
//	@Router			/biz/form							[post]
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
