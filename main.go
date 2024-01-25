package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello from Snippetbox"))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Display a specific snippet..."))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST"
	if r.Method != http.MethodPost {
		// w.Header().Set("Allow", "POST")
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// initialize a new servemux
	mux := http.NewServeMux()
	// register the home function as the handler url
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}

// 34page
