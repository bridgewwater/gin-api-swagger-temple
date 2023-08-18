## swaggo

- [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

## Attribute

```api
// @Param   enumstring  query     string     false  "string enums"       Enums(A, B, C)
    // @Param   enumint     query     int        false  "int enums"          Enums(1, 2, 3)
    // @Param   enumnumber  query     number     false  "int enums"          Enums(1.1, 1.2, 1.3)
    // @Param   string      query     string     false  "string valid"       minlength(5)  maxlength(10)
    // @Param   int         query     int        false  "int valid"          minimum(1)    maximum(10)
    // @Param   default     query     string     false  "string default"     default(A)
    // @Param   example     query     string     false  "string example"     example(string)
    // @Param   collection  query     []string   false  "string collection"  collectionFormat(multi)
    // @Param   extensions  query     []string   false  "string collection"  extensions(x-example=test,x-nullable)
```

It also works for the struct fields

```go
type Foo struct {
Bar string `minLength:"4" maxLength:"16" example:"random string"`
Baz int `minimum:"10" maximum:"20" default:"15"`
Qux []string `enums:"foo,bar,baz"`
}
```

### 当前可用的

| 字段名              | 类型        | 描述                                                                                                                                                                                                                                                                            |
|------------------|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| default          | *         | 声明如果未提供任何参数，则服务器将使用的默认参数值，例如，如果请求中的客户端未提供该参数，则用于控制每页结果数的“计数”可能默认为100。 （注意：“default”对于必需的参数没有意义）。参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-6.2。 与JSON模式不同，此值务必符合此参数的定义[类型](#parameterType)。                                                   |
| maximum          | `number`  | 参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.2.                                                                                                                                                                                             |
| minimum          | `number`  | 参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.3.                                                                                                                                                                                             |
| maxLength        | `integer` | 参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.1.                                                                                                                                                                                             |
| minLength        | `integer` | 参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.2.                                                                                                                                                                                             |
| enums            | [\*]      | 参看 https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.5.1.                                                                                                                                                                                             |
| format           | `string`  | 上面提到的[类型](#parameterType)的扩展格式。有关更多详细信息，请参见[数据类型格式](https://swagger.io/specification/v2/#dataTypeFormat)。                                                                                                                                                                     |
| collectionFormat | `string`  | 指定query数组参数的格式。 可能的值为： <ul><li>`csv` - 逗号分隔值 `foo,bar`. <li>`ssv` - 空格分隔值 `foo bar`. <li>`tsv` - 制表符分隔值 `foo\tbar`. <li>`pipes` - 管道符分隔值 <code>foo&#124;bar</code>. <li>`multi` - 对应于多个参数实例，而不是单个实例 `foo=bar＆foo=baz` 的多个值。这仅对“`query`”或“`formData`”中的参数有效。 </ul> 默认值是 `csv`。 |

### 进一步的

| 字段名         |    类型     | 描述                                                                                 |
|-------------|:---------:|------------------------------------------------------------------------------------|
| multipleOf  | `number`  | See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.1. |
| pattern     | `string`  | See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.3. |
| maxItems    | `integer` | See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.2. |
| minItems    | `integer` | See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.3. |
| uniqueItems | `boolean` | See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.4. |

## Mime类型

`swag` 接受所有格式正确的MIME类型, 即使匹配 `*/*`。除此之外，`swag`还接受某些MIME类型的别名，如下所示：

| Alias                 | MIME Type                         |
|-----------------------|-----------------------------------|
| json                  | application/json                  |
| xml                   | text/xml                          |
| plain                 | text/plain                        |
| html                  | text/html                         |
| mpfd                  | multipart/form-data               |
| x-www-form-urlencoded | application/x-www-form-urlencoded |
| json-api              | application/vnd.api+json          |
| json-stream           | application/x-json-stream         |
| octet-stream          | application/octet-stream          |
| png                   | image/png                         |
| jpeg                  | image/jpeg                        |
| gif                   | image/gif                         |


## 参数类型

- query
- path
- header
- body
- formData

## 数据类型

- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- user defined struct