package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current URL path exactly matches "/".
	// If it does not, use the http.NotFound() function to send a 404 response to the client
	// Importantly, we then return from the handler.
	// If we don't return, the handler would keep executing and also write the "Hello from Snippetbox!" message
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

	// w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// without using golang's http constants:
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	http.Error(w, "Method Not Allowed", 405)
	// 	// The above is equivalent to:
	// 	// w.WriteHeader(405)
	// 	// w.Write([]byte("Method Not Allowed"))
	// 	return
	// }

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
