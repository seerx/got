package gottp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Router 路由定义
// 对 httprouter.Router 包装
type Router struct {
	Debug          func(messsage string)
	Error          func(err error)
	router         *httprouter.Router
	beforeHandlers []Intercptor
	afterHandlers  []Intercptor
	panicHandlers  []PanicHandler
}

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
	return &Router{
		Debug:  func(messsage string) {},
		Error:  func(err error) {},
		router: httprouter.New(),
	}
}

//GetHTTPRouter 获取 http handler
func (rc *Router) GetHTTPRouter() *httprouter.Router {
	return rc.router
}

//AddBeforeInterceptor 添加执行前拦截器
func (rc *Router) AddBeforeInterceptor(interceptor Intercptor) {
	rc.beforeHandlers = append(rc.beforeHandlers, interceptor)
}

// AddAfterInterceptor 添加执行后拦截器，
// 请勿在 AfterInterceptor 函数中调用 ctx.ResponseHeaderSet 和 ctx.ResponseHeaderAdd 函数
func (rc *Router) AddAfterInterceptor(interceptor Intercptor) {
	rc.afterHandlers = append(rc.afterHandlers, interceptor)
}

//AddPanicHandler 添加错误处理函数
func (rc *Router) AddPanicHandler(handler PanicHandler) {
	rc.panicHandlers = append(rc.panicHandlers, handler)
}

func (rc *Router) forJumpout(path string, ctx *Context) {
	if err := recover(); err != nil {
		if _, ok := err.(JumpoutError); !ok {
			// 不是 JumpoutError 错误, 非正常结束

			// 执行错误处理函数
			throw := len(rc.panicHandlers) == 0
			for _, h := range rc.panicHandlers {
				handled := h(path, ctx, err)
				if !handled {
					// 如果没有处理
					throw = true
				}
			}
			if throw {
				panic(err)
			}
		} else {
			// 是 JumpoutError 错误,正常结束
			// 执行 after 拦截器链
			for _, i := range rc.afterHandlers {
				i(path, ctx)
			}
		}
	}
}

func (rc *Router) executeBefore(path string, ctx *Context) {
	for _, i := range rc.beforeHandlers {
		i(path, ctx)
	}
}

//GET 注册 GET 请求
func (rc *Router) GET(path string, handler Handler) {
	rc.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var o = Context{
			Request: r,
			Writer:  w,
			params:  params,
			router:  rc,
		}
		defer rc.forJumpout(path, &o)
		rc.executeBefore(path, &o)
		handler(&o)
	})
}

//POST 注册 POST 请求
func (rc *Router) POST(path string, handler Handler) {
	rc.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var o = Context{
			Request: r,
			Writer:  w,
			params:  params,
			router:  rc,
		}
		defer rc.forJumpout(path, &o)
		rc.executeBefore(path, &o)
		handler(&o)
	})
}

//PUT 注册 PUT 请求
func (rc *Router) PUT(path string, handler Handler) {
	rc.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var o = Context{
			Request: r,
			Writer:  w,
			params:  params,
			router:  rc,
		}
		defer rc.forJumpout(path, &o)
		rc.executeBefore(path, &o)
		handler(&o)
	})
}

//DELETE 注册 DELETE 请求
func (rc *Router) DELETE(path string, handler Handler) {
	rc.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		var o = Context{
			Request: r,
			Writer:  w,
			params:  params,
			router:  rc,
		}
		defer rc.forJumpout(path, &o)
		rc.executeBefore(path, &o)
		handler(&o)
	})
}

//HEAD 注册 HEAD 请求
func (rc *Router) HEAD(path string, handler Handler) {
	rc.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var o = Context{
			Request: r,
			Writer:  w,
			params:  params,
			router:  rc,
		}
		defer rc.forJumpout(path, &o)
		rc.executeBefore(path, &o)
		handler(&o)
	})
}
