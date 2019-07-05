package elfin

import (
	"os"
	"os/signal"

	elfin "github.com/obipawan/elfin/lifecycle"
)

/*
GracefulStop Helps avoiding boilerplate code and provides a nicer chain API to
perform your server cleanups gracefully.

Here's an example of how to use this
gracefulStop.New().
		Notify(
			syscall.SIGTERM,
			syscall.SIGINT,
			...
			...
		).Laters(
			func(){ // clean up code for server or db },
			func(){ // others },
			...
			...
		)
*/
type GracefulStop struct {
	c chan os.Signal
}

/*
NewGracefulStop returns a pointer to a new GracefulStop instance
*/
func NewGracefulStop() *GracefulStop {
	kelly := &GracefulStop{}
	kelly.c = make(chan os.Signal)
	return kelly
}

/*
Notify Similar to signal.Notify except creates it's own channel to read from
and returns a pointer to a GracefulStop instace
*/
func (kelly *GracefulStop) Notify(signals ...os.Signal) *GracefulStop {
	signal.Notify(kelly.c, signals...)
	return kelly
}

/*
Laters waits on a signal from flags subscribed to on GracefulStop.Notify and
executes these functions. You'll typically provide clean up functions here
*/
func (kelly *GracefulStop) Laters(funcs ...elfin.Func) {
	<-kelly.c
	for _, laterfunc := range funcs {
		laterfunc(nil)
	}
	os.Exit(0)
}
