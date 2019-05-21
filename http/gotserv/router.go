package gotserv

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Router 路由定义
// 对 httprouter.Router 包装
type Router struct {
	router *httprouter.Router
}

//Handler web服务函数定义
type Handler func(context *HTTPContext)

//Filter 过滤器
// func Filter(h Handler) RequestHandler {
// 	return func(o httputil.HTTPObject) {
// 		auth := o.GetQueryParam("auth")
// 		if auth == "" || auth != "0000123418838121" {
// 			o.ReturnStatus(1, "不可信的调用")
// 		} else {
// 			h(o)
// 		}
// 	}
// }

//NewRouter 创建新的路由
func NewRouter() *Router {
	return &Router{router: httprouter.New()}
}

//GetHTTPRouter 获取 http handler
func (rc *Router) GetHTTPRouter() *httprouter.Router {
	return rc.router
}

func forJumpout() {
	if err := recover(); err != nil {
		if _, ok := err.(JumpoutError); !ok {
			// 不是 JumpoutError 错误, 抛出
			panic(err)
		} else {
			// 是 JumpoutError 错误,不作任何操作
		}
	}
}

//GET 注册 GET 请求
func (rc *Router) GET(path string, handler Handler) {
	rc.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer forJumpout()
		var o = HTTPContext{
			Request: r,
			Writer:  w,
			params:  params,
		}
		handler(&o)
	})
}

//POST 注册 POST 请求
func (rc *Router) POST(path string, handler Handler) {
	rc.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer forJumpout()
		var o = HTTPContext{
			Request: r,
			Writer:  w,
			params:  params,
		}
		handler(&o)
	})
}

//PUT 注册 PUT 请求
func (rc *Router) PUT(path string, handler Handler) {
	rc.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer forJumpout()
		var o = HTTPContext{
			Request: r,
			Writer:  w,
			params:  params,
		}
		handler(&o)
	})
}

//DELETE 注册 DELETE 请求
func (rc *Router) DELETE(path string, handler Handler) {
	rc.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer forJumpout()
		var o = HTTPContext{
			Request: r,
			Writer:  w,
			params:  params,
		}
		handler(&o)
	})
}

//HEAD 注册 HEAD 请求
func (rc *Router) HEAD(path string, handler Handler) {
	rc.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer forJumpout()
		var o = HTTPContext{
			Request: r,
			Writer:  w,
			params:  params,
		}
		handler(&o)
	})
}
