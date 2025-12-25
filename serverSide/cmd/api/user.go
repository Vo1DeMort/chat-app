package main

import "net/http"

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request)    {}

// update users' password, email, or role
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {}

// get list of users based on roles or project access
func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {}
func (app *application) getUser(w http.ResponseWriter, r *http.Request)  {}
