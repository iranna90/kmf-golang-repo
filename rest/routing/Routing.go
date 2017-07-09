package routing

import (
	"net/http"
	"regexp"
	"log"
	"strings"
	"fmt"
	"sync/atomic"
)

type MyRouter struct {
	Routes map[*regexp.Regexp]Routing
}

type PathParams map[string]string

type Handler func(http.ResponseWriter, *http.Request)

type HandlerWithParam func(w http.ResponseWriter, r *http.Request, params *PathParams)

const (
	exclude                            int    = 1
	opening, closing, pathParamMatcher string = "{", "}", ".*"
)

type PathParamsDetails map[uint8]paramDetails

type paramDetails struct {
	paramName    string
	locationFrom uint8
}

type Routing interface {
	getHandler() Handler

	getHandlerWithParam() HandlerWithParam

	getPathParameters() PathParamsDetails

	getAction() string
}

type Route struct {
	method     string
	Path       string
	consumes   string
	allowedFor string
}

type RouteWithoutParam struct {
	Route
	Handler Handler
}

type RouteWithParam struct {
	Route
	pathParameters PathParamsDetails
	Handler        HandlerWithParam
}

func (r RouteWithoutParam) getHandler() Handler {
	return r.Handler
}

func (r RouteWithoutParam) getAction() string {
	return r.method
}

func (r RouteWithoutParam) getPathParameters() PathParamsDetails {
	return nil
}

func (r RouteWithoutParam) getHandlerWithParam() HandlerWithParam {
	panic("Un supported for url without path param")
}

func (router *MyRouter) Register(route RouteWithoutParam) {
	path := route.Path

	pattern, err := regexp.Compile(path)
	if err != nil {
		log.Panicf("Invalid url pattern %s", path)
		panic(path)
	}

	router.Routes[pattern] = route
}

func (router *MyRouter) RegisterWithParams(route RouteWithParam) {
	path := route.Path
	pathParams := PathParamsDetails{}
	var count uint32
	readPathParams(path, pathParams, &count)
	if len(pathParams) == 0 {
		log.Panicf("Invalid url: %s has no path parameters", path)
		panic(path)
	}

	pattern, err := regexp.Compile(buildPattern(path, pathParams))

	if err != nil {
		log.Panicf("Error while prepairing pattern %s", path)
		panic(path)
	}
	route.pathParameters = pathParams
	router.Routes[pattern] = route
}

func (r RouteWithParam) getHandler() Handler {
	panic("Un supported for url without path param")
}

func (r RouteWithParam) getHandlerWithParam() HandlerWithParam {
	return r.Handler
}

func (r RouteWithParam) getAction() string {
	return r.method
}

func (r RouteWithParam) getPathParameters() PathParamsDetails {
	return r.pathParameters
}

func buildPattern(path string, paramsDetails PathParamsDetails) string {
	for _, v := range paramsDetails {
		format := getOldFormat(v.paramName)
		path = strings.Replace(path, format, pathParamMatcher, 1)
	}
	return path
}

func getOldFormat(pathParamName string) string {
	return opening + pathParamName + closing
}

func readPathParams(path string, params PathParamsDetails, count *uint32) {
	index := strings.Index(path, opening)
	if index == -1 {
		return
	}

	closingIndex := strings.Index(path, closing)
	if closingIndex == -1 {
		panic(fmt.Sprintf("Invalid url: %s does not contains closing brace ", path))
	}

	pathParam := path[index+exclude:closingIndex]
	atomic.AddUint32(count, uint32(exclude))
	params[uint8(*count)] = paramDetails{pathParam, uint8(index)}

	readPathParams(path[closingIndex+exclude:], params, count)
}

func populatePathParams(path string, paramDetails PathParamsDetails, params PathParams, paramNumber *uint32) {

	if *paramNumber > uint32(len(paramDetails)) {
		return
	}

	details := paramDetails[uint8(*paramNumber)]
	path = path[details.locationFrom:]
	closingIndex := strings.Index(path, "/");
	var value string
	if closingIndex != -1 {
		value = path[:closingIndex]
	} else {
		value = path
	}

	params[details.paramName] = value
	if closingIndex != -1 {
		path = path[closingIndex:]
		atomic.AddUint32(paramNumber, 1)
		populatePathParams(path, paramDetails, params, paramNumber)
	}
}

func (router *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	headers := r.Header
	fmt.Println("Headers are ", headers)
	w.Header().Add("owner", "Iranna Go-Lang")
	w.Header().Add("description", "learning go-lang first phase")
	var matched bool
	for pattern, route := range router.Routes {
		if pattern.MatchString(path) {
			matched = true
			// without param
			if route.getPathParameters() == nil {
				route.getHandler()(w, r)
			} else {
				pathParams := PathParams{}
				var count uint32 = 1
				populatePathParams(path, route.getPathParameters(), pathParams, &count)
				route.getHandlerWithParam()(w, r, &pathParams)
			}
		}
	}
	if !matched {
		w.WriteHeader(404)
		w.Write([]byte("NOT FOUND≈≈≈"))
	}
}
