package elfin

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	elfin "github.com/obipawan/elfin/lifecycle"
	"github.com/obipawan/elfin/middlewares"
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
	addr        string
}

/*
New returns a new instance
*/
func New() *Elfin {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	elfin := &Elfin{addr: os.Getenv("HOST") + ":" + port}
	elfin.Mux = NewRouter().Mux
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
	server := &http.Server{
		Addr:    elfin.addr,
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
		Notify(syscall.SIGTERM, syscall.SIGINT).
		Laters(elfin.OnShutdownFuncs...)

	elfin.handleOnPostStart()

	if err := server.ListenAndServe(); err != nil {
		if elfin.CanReload(err, *elfin.GetReloadOptions().OnStartError) {
			elfin.handleOnReload(err)
			return
		}
		panic(err)
	}
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
