package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// custom error routes
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// doc
	router.HandlerFunc(http.MethodGet, "/v1/docs", app.docs)

	// user
	router.HandlerFunc(http.MethodPost, "/v1/register", app.registerUser)
	router.HandlerFunc(http.MethodPost, "/v1/login", app.loginUser)
	router.HandlerFunc(http.MethodPut, "/v1/update-user", app.updateUser)
	router.HandlerFunc(http.MethodGet, "/v1/get-users", app.getUsers)
	router.HandlerFunc(http.MethodGet, "/v1/get-users/id", app.getUser)

	//todos
	router.HandlerFunc(http.MethodGet, "/v1/get-todos", app.getTodos)
	router.HandlerFunc(http.MethodGet, "/v1/get-todos/id", app.getTodo)
	router.HandlerFunc(http.MethodPost, "/v1/post-todo", app.postTodo)
	router.HandlerFunc(http.MethodDelete, "/v1/delet-todo", app.deleteTodo)

	// projects
	router.HandlerFunc(http.MethodPost, "/v1/create-project", app.createProject)
	router.HandlerFunc(http.MethodPut, "/v1/update-project", app.updateProject)
	router.HandlerFunc(http.MethodDelete, "/v1/delet-project", app.deleteProject)
	router.HandlerFunc(http.MethodPost, "/v1/assign-task", app.assignTasks)
	router.HandlerFunc(http.MethodGet, "/v1/get-projects", app.getProject)
	router.HandlerFunc(http.MethodGet, "/v1/get-projects/id", app.getProject)

	// every route and revocer from panic
	// NOTE: the middleware can only revocer from this go routine
	return app.recoverPanic(router)
}
