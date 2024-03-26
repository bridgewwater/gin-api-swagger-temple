package biz

import "github.com/gin-gonic/gin"

// FileUpload
//
//	@Summary		Upload file for count file size
//	@Description	Upload file. warning api in prod will hide, abs remote api for dev
//	@Tags			biz
//	@ID				file.upload
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			file	formData	file			true	"this is a test file"
//
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	errdef.Err		"We need ID!!"
//	@Failure		404		{object}	errdef.Err		"Can not find ID"
//	@Router			/file/upload						[post]
func FileUpload(c *gin.Context) {

}
