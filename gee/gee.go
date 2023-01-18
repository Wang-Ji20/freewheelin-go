package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	e := group.engine
	newG := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: e,
	}
	e.groups = append(e.groups, newG)
	return newG
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.add(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc)  {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc)  {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
