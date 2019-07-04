package elfin

const (
	// ShouldPanic option kills the server in case of an error
	ShouldPanic = iota
	// ShouldReload option attempts reloads the server in case of an error
	ShouldReload
)

/*
ReloadOptions represents all possible reload options to elfin
*/
type ReloadOptions struct {
	OnStartError    *int
	OnPreStartError *int
}

/*
Options returns a pointer to a new ReloadOption
*/
func Options() *ReloadOptions {
	return &ReloadOptions{}
}

/*
SetOnStartError specifies what should happen onStartError
*/
func (ro *ReloadOptions) SetOnStartError(option int) *ReloadOptions {
	ro.OnStartError = &option
	return ro
}

/*
SetOnPreStartError specifies what should happen onPreStartError
*/
func (ro *ReloadOptions) SetOnPreStartError(option int) *ReloadOptions {
	ro.OnPreStartError = &option
	return ro
}
