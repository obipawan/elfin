package middlewares

import (
	"fmt"
	"net/http"
)

/*
Log middleware logs requests
*/
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
