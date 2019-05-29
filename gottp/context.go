package gottp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/seerx/got/cache"
)

//Context http 请求信息定义
type Context struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	params      httprouter.Params
	queryValues url.Values
	router      *Router
	session     cache.Entity
}

// Intercptor 拦截器定义
// 如果要不执行后续操作，直接调用 ctx.Return* 函数
type Intercptor func(path string, ctx *Context)

// PanicHandler 错误处理函数定义
// return bool
//			true	错误已经被处理，不再抛出
//			false	错误未被处理，抛出
type PanicHandler func(path string, context *Context, err interface{}) bool

// Handler web服务函数定义
type Handler func(context *Context)

//ResponseHeader json 请求返回的基础信息
// type ResponseHeader struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"msg"`
// }

//Response json 请求返回带有数据的信息
// type Response struct {
// 	ResponseHeader
// 	Data interface{} `json:"data"`
// }

//JumpoutError 跳过后续代码用到的错误定义
type JumpoutError struct {
	err string
}

func (e *JumpoutError) Error() string {
	return e.err
}

//ResponseHeaderSet 设置返回头
func (o *Context) ResponseHeaderSet(key, value string) {
	o.Writer.Header().Set(key, value)
}

//ResponseHeaderAdd 追加返回头
func (o *Context) ResponseHeaderAdd(key, value string) {
	o.Writer.Header().Set(key, value)
}

//DecodeRequestBodyAsJSON 把请求内容解析为 json
func (o *Context) DecodeRequestBodyAsJSON(v interface{}) error {
	return json.NewDecoder(o.Request.Body).Decode(v)
}

//ResponseText 返回文本
func (o *Context) ResponseText(text string) {
	o.ResponseTextf(200, text)
}

//ResponseTextStatus 返回数据并指定状态
func (o *Context) ResponseTextStatus(status int, text string) {
	o.ResponseTextf(status, text)
}

//ResponseTextf 返回 string
func (o *Context) ResponseTextf(status int, formatter string, a ...interface{}) {
	msg := fmt.Sprintf(formatter, a...)
	data := []byte(msg)
	o.ResponseHeaderSet("Content-Type", "text/plain; charset=utf-8")
	o.Writer.WriteHeader(status)
	o.Writer.Write(data)
	panic(JumpoutError{"jump-out"})
}

//ResponseJSONStatus 返回 JSON 对象
// 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
func (o *Context) ResponseJSONStatus(status int, jsonObject interface{}) error {
	data, err := json.Marshal(jsonObject)
	if err == nil {
		o.ResponseHeaderSet("Content-Type", "application/json; charset=utf-8")
		o.Writer.WriteHeader(status)

		o.Writer.Write(data)
		panic(JumpoutError{"jump-out"})
	}

	return err
}

//ResponseJSON 返回 json 数据
func (o *Context) ResponseJSON(jsonObject interface{}) error {
	return o.ResponseJSONStatus(200, jsonObject)
}

//ReturnHeader 返回状态
// 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
// func (o *Context) ReturnHeader(code int, formatter string, a ...interface{}) error {
// 	msg := fmt.Sprintf(formatter, a...)
// 	return o.ReturnJSON(ResponseHeader{code, msg})
// }

// //ReturnData 返回数据
// // 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
// func (o *Context) ReturnData(code int, message string, data interface{}) error {
// 	var res = Response{
// 		ResponseHeader: ResponseHeader{Code: code, Message: message},
// 		Data:           data,
// 	}

// 	return o.ReturnJSON(res)
// }

//ReturnFile 返回文件，执行该函数不会跳过后面的代码
func (o *Context) ReturnFile(filePath string) {
	http.ServeFile(o.Writer, o.Request, filePath)
}

//ParamInURL 获取 url 中名称为 name 的第一个参数的值
// 注意是 ? 之后的 name=value 形式的参数
func (o *Context) ParamInURL(name string) string {
	if o.queryValues == nil {
		o.queryValues = o.Request.URL.Query()
	}
	return o.queryValues.Get(name)
}

//NamedParamInURL 获取 url 中命名的参数，参见 httprouter 的参数名称
// 例如:   /getinfo/:user
//	获取 :user 位置的参数，可使用此函数
//		NamedParamInURL('user')
func (o *Context) NamedParamInURL(name string) string {
	return o.params.ByName(name)
}

//ParamInForm 获取 form 中提交的参数
func (o *Context) ParamInForm(name string) string {
	return o.Request.PostFormValue(name)
}

//GetSeesion 获取 session
// 使用 session 的话，需要在 main 包中初始化
// 例：
// 	cache := cache.NewCacheManager(memcache.PROVIDER, 600)
// 	session.Init("go-session", cache)
func (o *Context) GetSeesion() cache.Entity {
	if o.session == nil {
		if SSManager == nil {
			panic(fmt.Errorf("You need init session manager before use it, Call session.Init in main package's init func"))
		}

		entity := SSManager.SessionStart(o.Writer, o.Request)
		o.session = entity
	}
	return o.session
}
