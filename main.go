package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	fmt.Fprintf(w, "Hello, %s!\n", user)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{user}", HelloHandler)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
