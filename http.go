package got

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/seerx/goql"
	"github.com/seerx/got/gottp"
	"github.com/seerx/gql"
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

// DEBUG 设置 debug 函数
func DEBUG(fn func(string)) {
	checkInit()
	def.Debug = fn
}

// ERROR 设置 错误信息 函数
func ERROR(fn func(error)) {
	checkInit()
	def.Error = fn
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

func gqlServe(context *gottp.Context, gqlHandler http.Handler) {
	gqlHandler.ServeHTTP(context.Writer, context.Request)
}

// InitGraphQL 初始化 GraphQL
func InitGraphQL(g *gql.GQL, path string, cfg *handler.Config) {
	var gqlHandler = g.NewHandler(cfg)
	def.GET(path, func(context *gottp.Context) {
		gqlServe(context, gqlHandler)
	})
	def.POST(path, func(context *gottp.Context) {
		gqlServe(context, gqlHandler)
	})
}

// InitGoQL 初始化 goql
func InitGoQL(g *goql.GQL, path string, cfg *handler.Config) {
	handler := g.CreateHandler(cfg)
	def.GET(path, func(context *gottp.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	})
	def.POST(path, func(context *gottp.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	})
}
