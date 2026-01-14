package main

import "net/http"

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	// check if the user already exist with the provided credentials
	// if registration success , call login
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// generate token ( short live)
	// check the token in the auth middleware
	// keep the refresh token in the db ( long live )
	// need a func to validate the refersh token and generate new one
}

// update users' password, email, or role
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {}

// get list of users based on roles or project access
func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {}
func (app *application) getUser(w http.ResponseWriter, r *http.Request)  {}
