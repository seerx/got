package got

import (
	"fmt"

	"github.com/seerx/got/gottp"
)

var def *gottp.Router

func init() {
	def = gottp.NewRouter()
}

func checkInit() {
	if def == nil {
		panic(fmt.Errorf("请在 package main 中 import \"exthttp\""))
	}
}

//DefaultRouter 获取默认的 Router
func DefaultRouter() *gottp.Router {
	checkInit()
	return def
}

//GET 在默认的 defaultConetx 上注册 GET 请求
func GET(path string, handler gottp.Handler) {
	checkInit()
	def.GET(path, handler)
}

//POST 在默认的 defaultConetx 上注册 POST 请求
func POST(path string, handler gottp.Handler) {
	checkInit()
	def.POST(path, handler)
}

//HEAD 在默认的 defaultConetx 上注册 HEAD 请求
func HEAD(path string, handler gottp.Handler) {
	checkInit()
	def.HEAD(path, handler)
}

//PUT 在默认的 defaultConetx 上注册 PUT 请求
func PUT(path string, handler gottp.Handler) {
	checkInit()
	def.PUT(path, handler)
}

//DELETE 在默认的 defaultConetx 上注册 DELETE 请求
func DELETE(path string, handler gottp.Handler) {
	checkInit()
	def.DELETE(path, handler)
}
