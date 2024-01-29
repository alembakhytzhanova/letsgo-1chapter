package main

import "net/http"

func (app *applicaton) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.SnippetCreate)
	return mux
}
