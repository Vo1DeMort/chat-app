package data

import (
	"database/sql"
	"log"
	"time"
)

type Project struct {
	Id              int64 // uuid might be better
	Name            string
	Created         time.Time
	PeopeWithAccess []User
	Todos           []Todos
}

type ProjectModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m ProjectModel) Create(p *Project) error {
	return nil
}

func (m ProjectModel) Update(p *Project) error {
	return nil
}

// admin access
func (m ProjectModel) Delete(p *Project) error {
	return nil
}
func (m ProjectModel) AssignTasks(p *Project) error {
	return nil
}
func (m ProjectModel) GetProjects(p *Project) error {
	return nil
}
func (m ProjectModel) GetProject(p *Project) error {
	return nil
}
