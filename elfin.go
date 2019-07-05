package elfin

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

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

/*
Start starts the server
*/
func (elfin *Elfin) Start() {
	elfin.handleStart()
}

func (elfin *Elfin) handleStart() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := &http.Server{
		Addr:    os.Getenv("HOST") + ":" + port,
		Handler: chain(elfin.Mux, elfin.middlewares...),
	}

	elfin.OnShutdownFuncs = append(
		elfin.OnShutdownFuncs,
		func(interface{}) (error, []interface{}) {
			cx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			return server.Shutdown(cx), nil
		},
	)
	go NewGracefulStop().
		Notify(syscall.SIGTERM, syscall.SIGINT).
		Laters(elfin.OnShutdownFuncs...)

	server.ListenAndServe()
}
