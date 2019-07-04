package elfin

import (
	"net/http"
	"os"

	elfin "github.com/obipawan/elfin/lifecycle"
)

/*
Elfin .
*/
type Elfin struct {
	elfin.Lifecycle
	elfin.Reload
	Router
	Params
	middlewares []func(http.Handler) http.Handler
}

/*
New returns a new instance
*/
func New() *Elfin {
	elfin := &Elfin{}
	elfin.Mux = NewRouter().Mux
	return elfin
}

/*
ReloadOptions sets the reload options
*/
func (elfin *Elfin) ReloadOptions(options *elfin.ReloadOptions) {
	elfin.SetOptions(options)
}

/*
Start starts the server
*/
func (elfin *Elfin) Start() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	http.ListenAndServe(
		os.Getenv("HOST")+":"+port,
		chain(elfin.Mux, elfin.middlewares...),
	)
}

/*
Use appends a middleware. Middlewares are invoked in the order they're appended
*/
func (elfin *Elfin) Use(handler func(http.Handler) http.Handler) *Elfin {
	elfin.middlewares = append(elfin.middlewares, handler)
	return elfin
}

// Middleware chain
func chain(
	h http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	if len(middlewares) < 1 {
		return h
	}
	wrapped := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}
