package data

import (
	"database/sql"
	"log"
	"time"
)

type UserRole string

const (
	Developer UserRole = "Developer"
	Admin     UserRole = "ADMIN"
	Designer  UserRole = "Designer"
	Manager   UserRole = "Manager"
	Intern    UserRole = "Intern"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password password
	Role     UserRole
	Created  time.Time
}

type password struct {
	plaintext *string
	hash      []byte
}

// this for interacting with db
type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m UserModel) GetUsers(uer *User) error {
	return nil
}
func (m UserModel) GetUser(uer *User) error {
	return nil
}

// register
func (m UserModel) Create(uer *User) error {
	return nil
}

// update email, or password
func (m UserModel) Update(uer *User) error {
	return nil
}
func (m UserModel) Delete(uer *User) error {
	return nil
}
