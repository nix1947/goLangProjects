package main

import (
	"fmt"
	"log"
	"net/http"
)

// Two parameter
// w: HttpResponseWriter
// r: Request

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not support", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() error %v", err)
		return
	}

	fmt.Fprint(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s\n", address)
}

func main() {
	const rootDir = http.Dir("./static")
	fileserver := http.FileServer(rootDir)

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is listening on port 8000")

	var error = http.ListenAndServe("127.0.0.1:8000", nil)

	if error != nil {
		log.Fatal(error)
	}
}
