/*
Package elfin is a simple framework for building applications or services for
the web.

package main

import (
	"fmt"
	"net/http"

	"github.com/obipawan/elfin"
	"github.com/obipawan/elfin/middlewares"
)

func foo(w http.ResponseWriter, r *http.Request, p elfin.Params) {
	fmt.Fprint(w, p)
}

func main() {
	elfin := elfin.New()
	elfin.Get("/:word", foo)
	elfin.Use(middlewares.Log)
	elfin.Start()
}
*/
package elfin

import (
	"context"
	"net/http"
	"syscall"
	"time"

	"github.com/golang/glog"
	lc "github.com/obipawan/elfin/lifecycle"
	"github.com/obipawan/elfin/middlewares"
)

/*
Elfin describes the service or web application which can be used to subscribe to
lifecycles, set middlewares and add http routes.
*/
type Elfin struct {
	lc.Lifecycle
	lc.Reload
	Router
	Params
	middlewares []func(http.Handler) http.Handler
	addr        string
}

/*
ServerOpts defines options to setup the http server
*/
type ServerOpts struct {
	Address string
}

/*
New returns a new instance
*/
func New(opts *ServerOpts) *Elfin {
	options := opts
	if options == nil {
		options = &ServerOpts{Address: ":3000"}
	}
	elfin := &Elfin{addr: options.Address}
	elfin.Mux = NewRouter().Mux
	elfin.SetReloadOptions(
		lc.Options().
			SetOnPreStartError(lc.ShouldReload).
			SetOnStartError(lc.ShouldReload),
	)
	return elfin
}

/*
Use appends a middleware. Middlewares are invoked in the order they're appended
*/
func (elfin *Elfin) Use(handler func(http.Handler) http.Handler) *Elfin {
	elfin.middlewares = append(elfin.middlewares, handler)
	return elfin
}

/*
Start starts the server
*/
func (elfin *Elfin) Start() {
	elfin.StartWithAddr(elfin.addr)
}

/*
StartWithAddr starts the server with the given address host:port
*/
func (elfin *Elfin) StartWithAddr(address string) {
	addr := address
	if len(addr) == 0 {
		addr = elfin.addr
	}
	server := &http.Server{
		Addr:    addr,
		Handler: middlewares.Chain(elfin.Mux, elfin.middlewares...),
	}

	if err := elfin.handleOnPreStart(); err != nil {
		if elfin.CanReload(err, *elfin.GetReloadOptions().OnPreStartError) {
			elfin.handleOnReload(err)
			return
		}
		panic(err)
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
		Notify(syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT).
		Laters(elfin.OnShutdownFuncs...)

	elfin.handleOnPostStart()

	if err := server.ListenAndServe(); err != nil {
		if elfin.CanReload(err, *elfin.GetReloadOptions().OnStartError) {
			elfin.handleOnReload(err)
			return
		}
		panic(err)
	}

	glog.Info("Server started and listening on port ", elfin.addr)
}

/*
handleOnPrestart takes care of invoking registered pre-start callbacks
*/
func (elfin *Elfin) handleOnPreStart() error {
	for _, fun := range elfin.OnPreStartFuncs {
		if err, _ := fun(elfin.Mux); err != nil {
			return err
		}
	}
	return nil
}

/*
handleOnPostStart takes care of invoking registered post-start callbacks
*/
func (elfin *Elfin) handleOnPostStart() {
	for _, fun := range elfin.OnPostStartFuncs {
		go fun(elfin.Mux)
	}
}

/*
handleOnReload takes care of invoking registered reload callbacks
*/
func (elfin *Elfin) handleOnReload(err error) {
	elfin.BumpReloadCount()
	for _, fun := range elfin.OnReloadFuncs {
		fun(err) //pass error to handler
	}
	elfin.Start()
}
