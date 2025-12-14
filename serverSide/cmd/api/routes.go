package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	//WARNING: what is the diff with normal handler
	router.HandlerFunc(http.MethodGet, "/v1/getTodos", a.getTodos)
	router.HandlerFunc(http.MethodGet, "/v1/getTodo", a.getTodo)
	router.HandlerFunc(http.MethodPost, "/v1/postTodo", a.postTodo)
	router.HandlerFunc(http.MethodDelete, "/v1/deleteTodo", a.deleteTodo)

	return router
}
