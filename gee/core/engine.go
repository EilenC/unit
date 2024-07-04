package core

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc //中间件
	engine      *Engine
}

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New is the constructor of gee.Engine
func New() *Engine {
	eg := &Engine{router: newRouter()}
	eg.RouterGroup = &RouterGroup{engine: eg}
	eg.groups = []*RouterGroup{eg.RouterGroup}
	return eg
}

// Group is defined to create a new RouterGroup
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	eg := g.engine
	newGroup := &RouterGroup{prefix: g.prefix + prefix, engine: eg}
	eg.groups = append(eg.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, fn HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, fn)
}

// GET defines the method to add GET request
func (g *RouterGroup) GET(pattern string, fn HandlerFunc) {
	g.addRoute(http.MethodGet, pattern, fn)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(pattern string, fn HandlerFunc) {
	g.addRoute(http.MethodPost, pattern, fn)
}

func (e *Engine) Run(addr string) error {
	fmt.Println("Gee Serve running " + addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(newContext(w, r))
}
