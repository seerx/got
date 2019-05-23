package gotserv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/seerx/got/http/session"
)

//HTTPContext http 请求信息定义
type HTTPContext struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	params      httprouter.Params
	queryValues url.Values
}

//ResponseStatus json 请求返回的基础信息
type ResponseStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//Response json 请求返回带有数据的信息
type Response struct {
	ResponseStatus
	Data interface{} `json:"data"`
}

//JumpoutError 跳过后续代码用到的错误定义
type JumpoutError struct {
	err string
}

func (e *JumpoutError) Error() string {
	return e.err
}

//ReturnJSON 返回 JSON 对象
// 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
func (o *HTTPContext) ReturnJSON(jsonObject interface{}) error {
	data, err := json.Marshal(jsonObject)
	if err == nil {
		o.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		o.Writer.Header().Set("Content-Length", strconv.Itoa(len(data)))
		o.Writer.Write(data)

		panic(JumpoutError{"jump-out"})
	}

	return err
}

//ReturnStatus 返回状态
// 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
func (o *HTTPContext) ReturnStatus(status int, formatter string, a ...interface{}) error {
	msg := fmt.Sprintf(formatter, a...)
	return o.ReturnJSON(ResponseStatus{status, msg})
}

//ReturnData 返回数据
// 注意：如果该函数执行成功，则会跳过排在该函数后面的代码
func (o *HTTPContext) ReturnData(status int, message string, data *interface{}) error {
	var res = Response{
		ResponseStatus: ResponseStatus{Status: status, Message: message},
		Data:           data,
	}

	return o.ReturnJSON(res)
}

//ReturnFile 返回文件，执行该函数不会跳过后面的代码
func (o *HTTPContext) ReturnFile(filePath string) {
	http.ServeFile(o.Writer, o.Request, filePath)
}

//ParamInURL 获取 url 中名称为 name 的第一个参数的值
// 注意是 ? 之后的 name=value 形式的参数
func (o *HTTPContext) ParamInURL(name string) string {
	if o.queryValues == nil {
		o.queryValues = o.Request.URL.Query()
	}
	return o.queryValues.Get(name)
}

//NamedParamInURL 获取 url 中命名的参数，参见 httprouter 的参数名称
// 例如:   /getinfo/:user
//	获取 :user 位置的参数，可使用此函数
//		NamedParamInURL('user')
func (o *HTTPContext) NamedParamInURL(name string) string {
	return o.params.ByName(name)
}

//ParamInForm 获取 form 中提交的参数
func (o *HTTPContext) ParamInForm(name string) string {
	return o.Request.PostFormValue(name)
}

//GetSeesion 获取 session
func (o *HTTPContext) GetSeesion() session.Session {
	if session.AppSession == nil {
		panic(fmt.Errorf("You need init session manager before use it, Call session.InitSession( ... ) in main package's init func"))
	}

	ss := session.AppSession.SessionStart(o.Writer, o.Request)

	return ss
}
