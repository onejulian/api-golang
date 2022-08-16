package main

import (
	"fmt"
	"log"
	"net/http"
	"processXML/process"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", process.ProcessXML).Methods("GET")

	fmt.Println("Run server: http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", mux))
}