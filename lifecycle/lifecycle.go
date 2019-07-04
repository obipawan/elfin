package elfin

/*
Lifecycle .
*/
type Lifecycle struct {
	onPreStartFuncs,
	onPostStartFuncs,
	onShutdownFuncs,
	onReloadFuncs []Func
}

/*
OnPreStart registers callbacks for server pre-start hook
*/
func (lifecycle *Lifecycle) OnPreStart(funcs ...Func) *Lifecycle {
	lifecycle.onPreStartFuncs = funcs
	return lifecycle
}

/*
OnPostStart registers callbacks for server post-start hook
*/
func (lifecycle *Lifecycle) OnPostStart(funcs ...Func) *Lifecycle {
	lifecycle.onPostStartFuncs = funcs
	return lifecycle
}

/*
OnShutdown registers callbacks for server shutdown hook
*/
func (lifecycle *Lifecycle) OnShutdown(funcs ...Func) *Lifecycle {
	lifecycle.onShutdownFuncs = funcs
	return lifecycle
}

/*
OnReload registers callbacks for server reloads
*/
func (lifecycle *Lifecycle) OnReload(funcs ...Func) *Lifecycle {
	lifecycle.onReloadFuncs = funcs
	return lifecycle
}
