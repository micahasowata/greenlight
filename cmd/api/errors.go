package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}
