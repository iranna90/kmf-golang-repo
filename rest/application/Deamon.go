package main

import (
	. "random-repo-golang/rest/routing"
	"regexp"
	"net/http"
	"fmt"

)

func main() {
	myRouting := &MyRouter{make(map[*regexp.Regexp]Routing)}
	myRouting.Register(RouteWithoutParam{Route: Route{Path: "/main/hello"}, Handler: handlerMain})
	myRouting.RegisterWithParams(RouteWithParam{Route: Route{Path: "/main/abc/{run}"}, Handler: handle})
	fmt.Println("Starting to serve requests")
	http.ListenAndServe(":1234", myRouting)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func handle(w http.ResponseWriter, r *http.Request, params *PathParams) {
	w.Write([]byte((*params)["run"]))
}
