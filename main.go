package main

import (
	"fmt"
	"log"
	"net/http"
	"processXML/process"

)

func main() {
	data := process.ProcessXML()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, data)
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}