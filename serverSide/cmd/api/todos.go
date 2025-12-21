package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vo1DeMort/chat-app/serverSide/internal/data"
)

func (app *application) docs(w http.ResponseWriter, r *http.Request) {
	/// just messing with the golang
	err := app.writeJson(w, http.StatusOK, envelope{
		"doc": "endpoint for api doc",
	}, nil)
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	data := data.Todos{
		Id:      1,
		Created: time.Now(),
		Updated: time.Now(),
		Title:   "first todo",
		Type:    "testing",
	}

	err := app.writeJson(w, http.StatusOK, envelope{
		"todos": data,
	}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) getTodo(w http.ResponseWriter, r *http.Request) {
	//WARNING:  this should work with an id
	fmt.Fprintf(w, "get todo ")
}
func (app *application) postTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post todos ")
}
func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete todos ")
}
