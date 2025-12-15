package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// doc
	router.HandlerFunc(http.MethodGet, "/v1/docs", app.docs)

	//todos
	router.HandlerFunc(http.MethodGet, "/v1/getTodos", app.getTodos)
	router.HandlerFunc(http.MethodGet, "/v1/getTodo", app.getTodo)
	router.HandlerFunc(http.MethodPost, "/v1/postTodo", app.postTodo)
	router.HandlerFunc(http.MethodDelete, "/v1/deletTodo", app.deleteTodo)

	return router
}
