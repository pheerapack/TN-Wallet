package main

import (
	"log"
	"mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	//router.HandleFunc("/helloworld", getHelloworld).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
