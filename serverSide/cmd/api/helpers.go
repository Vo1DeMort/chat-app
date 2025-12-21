package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

// this is just a consitency for preference of json format
type envelope map[string]any

func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//NOTE: this might not be necessary for frontend program
	js = append(js, '\n')

	// for k, v := range headers {
	// 	w.Header()[k] = v
	// }
	maps.Insert(w.Header(), maps.All(headers))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}
