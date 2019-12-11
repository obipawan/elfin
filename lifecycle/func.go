package elfin

/*
Func are callbacks to lifecycle hooks, like onShutdown
*/
type Func func(interface{}) ([]interface{}, error)
