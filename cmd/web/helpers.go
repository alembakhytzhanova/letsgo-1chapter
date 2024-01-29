package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *applicaton) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applicaton) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *applicaton) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
