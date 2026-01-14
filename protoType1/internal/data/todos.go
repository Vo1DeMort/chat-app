package data

import "time"

// NOTE:
// struct tag to controll the appearence of key in the response , don't necessarily needed them
// can change the type of reponse, omit and hide with struct tags
type Todos struct {
	Id      int64
	Created time.Time
	Updated time.Time
	Title   string
	Type    string
}
