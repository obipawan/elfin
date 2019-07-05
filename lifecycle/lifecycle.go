package elfin

/*
Lifecycle .
*/
type Lifecycle struct {
	OnPreStartFuncs,
	OnPostStartFuncs,
	OnShutdownFuncs,
	OnReloadFuncs []Func
}

/*
OnPreStart registers callbacks for server pre-start hook
*/
func (lifecycle *Lifecycle) OnPreStart(funcs ...Func) *Lifecycle {
	lifecycle.OnPreStartFuncs = funcs
	return lifecycle
}

/*
OnPostStart registers callbacks for server post-start hook
*/
func (lifecycle *Lifecycle) OnPostStart(funcs ...Func) *Lifecycle {
	lifecycle.OnPostStartFuncs = funcs
	return lifecycle
}

/*
OnShutdown registers callbacks for server shutdown hook
*/
func (lifecycle *Lifecycle) OnShutdown(funcs ...Func) *Lifecycle {
	lifecycle.OnShutdownFuncs = funcs
	return lifecycle
}

/*
OnReload registers callbacks for server reloads
*/
func (lifecycle *Lifecycle) OnReload(funcs ...Func) *Lifecycle {
	lifecycle.OnReloadFuncs = funcs
	return lifecycle
}
