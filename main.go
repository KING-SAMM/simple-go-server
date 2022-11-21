package main

import (
	"fmt"
	"log"
	"net/htt"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Printf(w, "The post was sent successully")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Printf(w, "Hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.Handlefunc("/form", formHandler)
	http.Handlefunc("/hello", helloHandler)

	PORT := 8000 

	fmt.Printf("Server starting at port %v \n", PORT)

	if err := http.ListenAndServe(":8000", nil); err !nil {
		log.Fatal(err)
	}
}