package elfin

const (
	//MaxReload maximum number of reloads that are allowed
	MaxReload int = 3
)

/*
Reload describes the reload options and handles reloading the service if needed
while booting up.
*/
type Reload struct {
	options     *ReloadOptions
	reloadCount int
}

/*
CanReload checks if the server can reload, or panics based on the ReloadOption
*/
func (reload *Reload) CanReload(err error, option int) bool {
	if option == ShouldPanic {
		panic(err)
	}
	return reload.reloadCount < MaxReload
}

/*
BumpReloadCount increments the reload count
*/
func (reload *Reload) BumpReloadCount() {
	reload.reloadCount++
}

/*
SetReloadOptions sets the reload options
*/
func (reload *Reload) SetReloadOptions(options *ReloadOptions) {
	reload.options = options
}

/*
GetReloadOptions returns the reload options
*/
func (reload *Reload) GetReloadOptions() *ReloadOptions {
	return reload.options
}
