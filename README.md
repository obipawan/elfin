# Elfin

A simple framework for building applications or services for the web.

### Getting started
Make sure your [GOPATH](https://golang.org/doc/code.html#GOPATH) is set. Use [Glide](https://github.com/Masterminds/glide), [Dep](https://github.com/golang/dep) or other dependency management tool to install elfin

```sh
dep ensure -add github.com/obipawan/elfin
```
```go
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
```

### WIP