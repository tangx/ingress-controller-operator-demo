package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tangx/ingress-operator/pkg/httpx"
)

func main() {
	r := mux.NewRouter()

	// 两条路由
	r.NewRoute().Host("foo.com").Path("/").Methods("GET")
	r.NewRoute().Host("foo.com").Path(`/user/{id:\d+}`).Methods(httpx.MethodAny()...)

	// 创建 matcher
	match := &mux.RouteMatch{}

	for _, mock := range []struct {
		Method string
		URL    string
		Wanted bool
	}{
		{
			Method: http.MethodGet,
			URL:    "http://foo.com/",
			Wanted: true,
		},
	} {
		// 创建一个 请求
		req, _ := http.NewRequest(mock.Method, mock.URL, nil)
		// 匹配路由
		ret := r.Match(req, match)
		// 匹配结果
		fmt.Printf("wanted: %t, real: %t\n", mock.Wanted, ret)
	}
}
