package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
PanicHandler handles panics and responds with a 500 internal server error
*/
func PanicHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "Internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
