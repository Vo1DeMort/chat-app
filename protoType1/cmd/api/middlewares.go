package main

import (
	"fmt"
	"net/http"
)

// send error msg to client side ,then recover the panic
func (app *application) recoverPanic(next http.Handler) http.Handler {

	handler := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
