package main

import (
	"fmt"
	"net/http"
	"time"
)

// TODO: gonna try a todo api first , then add one by  one features
func (app *application) greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
