package main

import (
	"log"
	"net/http"
	"mux"
)

func main() {

	router := mux.NewRouter()

	//router.HandleFunc("/helloworld", getHelloworld).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}




