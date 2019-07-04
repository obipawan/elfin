package elfin

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handle is a function that can be registered to a route to handle HTTP
// requests. Like http.HandlerFunc, but has a third parameter for the values of
// wildcards (variables).
type Handle func(http.ResponseWriter, *http.Request, Params)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

/*
Router is a dedicated wrapper for any 3rd party router, with the intent to
replace with other routers (if needed) or, of course, build one.
*/
type Router struct {
	Mux *httprouter.Router
}

/*
NewRouter returns a pointer to a new Router instance
*/
func NewRouter() *Router {
	return &Router{Mux: httprouter.New()}
}

func handleProxy(handler Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
		params := Params{}
		for _, param := range p {
			params = append(params, Param{Key: param.Key, Value: param.Value})
		}
		handler(rw, rq, params)
	}
}

/*
Get adds a "GET" HTTP method for the given path
*/
func (router *Router) Get(path string, handler Handle) *Router {
	router.Mux.GET(path, handleProxy(handler))
	return router
}

/*
Put adds a "PUT" HTTP method for the given path
*/
func (router *Router) Put(path string, handler Handle) *Router {
	router.Mux.PUT(path, handleProxy(handler))
	return router
}

/*
Post adds a "POST" HTTP method for the given path
*/
func (router *Router) Post(path string, handler Handle) *Router {
	router.Mux.POST(path, handleProxy(handler))
	return router
}

/*
Patch adds a "PATCH" HTTP method for the given path
*/
func (router *Router) Patch(path string, handler Handle) *Router {
	router.Mux.PATCH(path, handleProxy(handler))
	return router
}

/*
Delete adds a "DELETE" HTTP method for the given path
*/
func (router *Router) Delete(path string, handler Handle) *Router {
	router.Mux.DELETE(path, handleProxy(handler))
	return router
}
