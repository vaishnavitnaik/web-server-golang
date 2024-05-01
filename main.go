package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello Vaishnavi")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull \n")
	name := r.FormValue("Name")
	address := r.FormValue("Address")

	fmt.Fprintf(w, "name = %s \n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}

func main() {
	fmt.Println("GO Lang Server")
	// informing go lang that it should check out static folder
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("panic")
	}
}
