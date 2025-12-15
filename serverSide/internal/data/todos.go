package data

import "time"

// struct tag to controll the appearence of key in the response
type Todos struct {
	Id      int64
	Created time.Time
	Updated time.Time
	Title   string
	Type    string
}
