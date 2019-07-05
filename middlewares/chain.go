package middlewares

import "net/http"

// Chain clubs middlewares
func Chain(
	h http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	if len(middlewares) < 1 {
		return h
	}
	wrapped := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}
