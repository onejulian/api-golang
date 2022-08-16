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

	mux.HandleFunc("/api", process.ProcessXML).Methods("GET")

	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}