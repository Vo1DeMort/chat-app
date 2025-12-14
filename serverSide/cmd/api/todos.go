package main

import (
	"fmt"
	"net/http"
)

// TODO: gonna try a todo api first , then add one by  one features
func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hello": "there",
		"info":  "there are some todos",
	}

	err := app.writeJson(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) getTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get todo ")
}
func (app *application) postTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post todos ")
}
func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete todos ")
}
