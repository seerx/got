# 概述
Golang 常用功能工具库

MD5、目录相关操作、格式化时间

Web Server 一些操作的封装，简化开发，增强功能

增加 session 功能

# 获取包
<pre>
go get github.com/seerx/got
</pre>

# web server
## 简单实例 main.go
<pre>
package main

import (
	"log"
	"net/http"

	"github.com/seerx/got"
	"github.com/seerx/got/gottp"
)

var svr *http.Server

func init() {
	svr = &http.Server{Addr: ":8099"}

	got.GET("/hello", hello)
}

func hello(ctx *gottp.Context) {
	ctx.ResponseText("Hello Golang!")
}

func main() {
	def := got.DefaultRouter()
	svr.Handler = def.GetHTTPRouter()
	log.Fatal(svr.ListenAndServe())
}

</pre>

在终端运行
<pre>
go run main.go
</pre>
在浏览器中输入 http://localhost:8099/hello
可以看到运行结果

## 使用 session 功能
首先在 mann.go 的 init() 函数中添加
<pre>
cache := cache.NewCacheManager(memcache.PROVIDER, 600)
session.Init("go-session", cache)
</pre>

在 handlers 中即可以使用 ctx.GetSession() 来获取和使用 session 了
<pre>
ss := ctx.GetSession()
ss.Set("test", "OK")
val := ss.Get("test")
</pre>