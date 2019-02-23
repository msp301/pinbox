package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello World!\n"))
	return
}

// our main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
