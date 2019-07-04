package elfin

const (
	//MaxReload maximum number of reloads that are allowed
	MaxReload int = 3
)

/*
Reload ..
*/
type Reload struct {
	options     *ReloadOptions
	reloadCount int
}

/*
CanReload checks if the server can reload, or panics based on the ReloadOption
*/
func (reload *Reload) CanReload(err error, options ReloadOptions) bool {
	if *options.OnPreStartError == ShouldPanic ||
		*options.OnStartError == ShouldPanic {
		panic(err)
	}
	return reload.reloadCount < MaxReload
}

/*
SetOptions sets the reload options
*/
func (reload *Reload) SetOptions(options *ReloadOptions) {
	reload.options = options
}

/*
GetOptions returns the reload options
*/
func (reload *Reload) GetOptions() *ReloadOptions {
	return reload.options
}
