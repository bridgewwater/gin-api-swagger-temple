// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},"swagger":"2.0","info":{"description":"{{escape .Description}}","title":"{{.Title}}","termsOfService":"http://github.com/","contact":{"name":"API Support","url":"http://github.com/","email":"support@sinlov.cn"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"version":"{{.Version}}"},"host":"{{.Host}}","basePath":"{{.BasePath}}","paths":{"/biz/json":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/json","responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"500":{"description":"Internal Server Error"}}}},"/biz/modelBiz":{"post":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/modelBiz","parameters":[{"description":"body biz.Biz for post","name":"biz","in":"body","required":true,"schema":{"$ref":"#/definitions/biz.Biz"}}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/path/{some_id}":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/path","parameters":[{"type":"string","description":"some id to show","name":"some_id","in":"path","required":true}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/query/":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/query","parameters":[{"type":"integer","description":"Offset","name":"offset","in":"query","required":true},{"type":"integer","description":"limit","name":"limit","in":"query"}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/string":{"get":{"description":"get string of this api.","tags":["biz"],"summary":"/biz/string","responses":{"200":{"description":"OK"},"500":{"description":"Internal Server Error"}}}}},"definitions":{"biz.Biz":{"type":"object","properties":{"id":{"type":"string","example":"id123zqqeeadg24qasd"},"info":{"type":"string","example":"input info here"},"limit":{"type":"integer","example":10},"offset":{"type":"integer","example":0}}},"errdef.Err":{"type":"object","properties":{"code":{"type":"integer"},"msg":{"type":"string"}}}},"securityDefinitions":{"BasicAuth":{"type":"basic"},"WithToken":{"description":"Please set the token of the API, note that it starts with \"Bearer\"","type":"apiKey","name":"Authorization","in":"header"}},"externalDocs":{"description":"OpenAPI","url":"https://swagger.io/resources/open-api/"}}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gin-api-swagger-temple",
	Description:      "This is a sample server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
