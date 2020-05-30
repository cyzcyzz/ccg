package ccg

import (
	"net/http"
)

// 框架入口
type HandlerFunc func(*Context)

type Engine struct {
	//router map[string]HandlerFunc
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//key := req.Method + "_" + req.URL.Path
	//if handler, ok := e.router[key]; ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 Not Found: %s\n", req.URL)
	//}
	c := newContext(w, req)
	e.router.handle(c)
}

// handle 接口需要实现serveHTTP方法，不然就报错
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
