package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	msg := `{"status" : "available", "environment" : %q, "version" : %q}`
	msg = fmt.Sprintf(msg, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(msg))
}
