package main

import (
	"encoding/json"
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

// / will do json paring ,validation
func (app *application) postTodo(w http.ResponseWriter, r *http.Request) {

	var input struct {
		// these fields must be in capital so that encode/json can see
		Title string `json:"title"`
		Type  string `json:"type"`
	}

	// must use non nil pointer value for Decode
	// otherwise json.invalid unmarsha error
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%v\n", input)
}
func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete todos ")
}
