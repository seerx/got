package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seerx/got"
	"github.com/seerx/got/pkg/gottp"
)

func main() {

	router := got.DefaultRouter()
	svr := &http.Server{Addr: fmt.Sprintf(":%d", 8080)}
	svr.Handler = router.GetHTTPRouter()

	got.GET("/hello", func(context *gottp.Context) {
		context.ResponseText("Hello got!")
	})

	log.Fatal(svr.ListenAndServe())
}
