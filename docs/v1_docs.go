// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplatev1 = `{
    "schemes": {{ marshal .Schemes }},"swagger":"2.0","info":{"description":"{{escape .Description}}","title":"{{.Title}}","termsOfService":"http://github.com/","contact":{"name":"API Support","url":"http://github.com/","email":"support@sinlov.cn"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"version":"{{.Version}}"},"host":"{{.Host}}","basePath":"{{.BasePath}}","paths":{"/biz/form":{"post":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/x-www-form-urlencoded"],"produces":["application/json"],"tags":["biz"],"summary":"api demo form with: x-www-form-urlencoded","parameters":[{"type":"string","default":"id123zqqeeadg24qasd","example":"id123zqqeeadg24qasd","name":"id","in":"formData"},{"type":"string","default":"info","example":"input info here","name":"info","in":"formData"},{"type":"integer","default":10,"example":10,"name":"limit","in":"formData"},{"type":"integer","default":0,"example":0,"name":"offset","in":"formData"}],"responses":{"200":{"description":"value in model.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/form_full":{"post":{"description":"for biz","consumes":["application/x-www-form-urlencoded"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/form_full","parameters":[{"type":"string","default":"foo","description":"form item foo","name":"foo","in":"formData","required":true},{"type":"string","default":"bar","description":"form item bar","name":"bar","in":"formData","required":true},{"type":"string","default":"baz","description":"form item baz","name":"baz","in":"formData","required":true}],"responses":{"200":{"description":"value in parse_http.FormContent","schema":{"$ref":"#/definitions/parse_http.FormContent"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/header_full":{"get":{"description":"for biz","consumes":["text/plain"],"produces":["text/plain"],"tags":["biz"],"summary":"/biz/header_full","parameters":[{"type":"string","default":"foo","description":"header BIZ_FOO ","name":"BIZ_FOO","in":"header","required":true},{"type":"string","default":"bar","description":"header BIZ_BAR ","name":"BIZ_BAR","in":"header","required":true}],"responses":{"200":{"description":"value in parse_http.HeaderContent","schema":{"$ref":"#/definitions/parse_http.HeaderContent"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/json":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"api demo json","responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"500":{"description":""}}}},"/biz/modelBiz":{"post":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"api json struct biz.Biz","parameters":[{"description":"body biz.Biz for post","name":"biz","in":"body","required":true,"schema":{"$ref":"#/definitions/biz.Biz"}}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/modelBizQuery":{"post":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"api post query with json struct biz.Biz","parameters":[{"type":"integer","description":"Offset","name":"offset","in":"query","required":true},{"type":"integer","description":"limit","name":"limit","in":"query"},{"description":"body biz.Biz for post","name":"biz","in":"body","required":true,"schema":{"$ref":"#/definitions/biz.Biz"}}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/path/{some_id}":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["multipart/form-data"],"produces":["application/json"],"tags":["biz"],"summary":"api path demo for route path","parameters":[{"type":"string","description":"some id to show","name":"some_id","in":"path","required":true}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/query":{"get":{"description":"warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"api demo for query.","parameters":[{"minimum":0,"type":"integer","default":0,"description":"Offset","name":"offset","in":"query","required":true},{"type":"integer","default":10,"description":"limit","name":"limit","in":"query"}],"responses":{"200":{"description":"value in biz.Biz","schema":{"$ref":"#/definitions/biz.Biz"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/query_full":{"get":{"description":"for biz","consumes":["application/json"],"produces":["application/json"],"tags":["biz"],"summary":"/biz/query_full","parameters":[{"type":"string","default":"\"\"","description":"params foo ","name":"foo","in":"query","required":true},{"type":"string","default":"\"\"","description":"params bar ","name":"bar","in":"query","required":true},{"type":"string","default":"\"\"","description":"params baz ","name":"baz","in":"query","required":true}],"responses":{"200":{"description":"value in parse_http.QueryContent","schema":{"$ref":"#/definitions/parse_http.QueryContent"}},"400":{"description":"error at errdef.Err","schema":{"$ref":"#/definitions/errdef.Err"}}}}},"/biz/string":{"get":{"description":"get string of this api. warning api in prod will hide, abs remote api for dev","consumes":["application/json"],"produces":["text/plain"],"tags":["biz"],"summary":"sample demo string","responses":{"200":{"description":"OK"},"500":{"description":""}}}},"/file/upload":{"post":{"description":"Upload file. warning api in prod will hide, abs remote api for dev","consumes":["multipart/form-data"],"produces":["application/json"],"tags":["biz"],"summary":"Upload file for count file size","operationId":"file.upload","parameters":[{"type":"file","description":"this is a test file","name":"file","in":"formData","required":true}],"responses":{"200":{"description":"ok","schema":{"type":"string"}},"400":{"description":"We need ID!!","schema":{"$ref":"#/definitions/errdef.Err"}},"404":{"description":"Can not find ID","schema":{"$ref":"#/definitions/errdef.Err"}}}}}},"definitions":{"biz.Biz":{"type":"object","properties":{"id":{"type":"string","default":"id123zqqeeadg24qasd","example":"id123zqqeeadg24qasd"},"info":{"type":"string","default":"info","example":"input info here"},"limit":{"type":"integer","default":10,"example":10},"offset":{"type":"integer","default":0,"example":0}}},"errdef.Err":{"type":"object","properties":{"code":{"type":"integer"},"msg":{"type":"string"}}},"parse_http.FormContent":{"type":"object","properties":{"headers":{"type":"object","additionalProperties":{"type":"string"}},"host":{"type":"string"},"method":{"type":"string"},"param":{"type":"string"},"parse_time":{"type":"string"},"postData":{"type":"object","additionalProperties":{"type":"string"}},"proto":{"type":"string"},"proto_major":{"type":"integer"},"proto_minor":{"type":"integer"},"queryString":{"type":"object","additionalProperties":{"type":"string"}},"referer":{"type":"string"},"remote_addr":{"type":"string"},"url":{"type":"string"}}},"parse_http.HeaderContent":{"type":"object","properties":{"header":{"type":"string"},"headers":{"type":"object","additionalProperties":{"type":"string"}},"host":{"type":"string"},"method":{"type":"string"},"param":{"type":"string"},"parse_time":{"type":"string"},"proto":{"type":"string"},"proto_major":{"type":"integer"},"proto_minor":{"type":"integer"},"referer":{"type":"string"},"remote_addr":{"type":"string"},"url":{"type":"string"}}},"parse_http.QueryContent":{"type":"object","properties":{"headers":{"type":"object","additionalProperties":{"type":"string"}},"host":{"type":"string"},"method":{"type":"string"},"param":{"type":"string"},"parse_time":{"type":"string"},"proto":{"type":"string"},"proto_major":{"type":"integer"},"proto_minor":{"type":"integer"},"queryString":{"type":"object","additionalProperties":{"type":"string"}},"referer":{"type":"string"},"remote_addr":{"type":"string"},"url":{"type":"string"}}}},"securityDefinitions":{"BasicAuth":{"type":"basic"},"WithToken":{"description":"Please set the token of the API, note that it starts with \"Bearer\"","type":"apiKey","name":"Authorization","in":"header"}},"externalDocs":{"description":"OpenAPI","url":"https://swagger.io/resources/open-api/"}}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "v1.x.x",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gin-api-swagger-temple",
	Description:      "This is a sample server",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}
