package main

import "net/http"

func (app *application) createProject(w http.ResponseWriter, r *http.Request) {}
func (app *application) updateProject(w http.ResponseWriter, r *http.Request) {}

// update users' password, email, or role
func (app *application) deleteProject(w http.ResponseWriter, r *http.Request) {}
func (app *application) assignTasks(w http.ResponseWriter, r *http.Request)   {}

// get list of users based on roles or project access
func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {}
func (app *application) getProject(w http.ResponseWriter, r *http.Request)  {}
