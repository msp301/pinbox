package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getLabels(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(`[ { "id": 765, "name": "Beep" } ]`))
}

func getMessages(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(`[ { "id": 999, "epoch": 1550930545657, "recipient": "Test <hello@test.com>", "sender": "API <api@server.net>", "subject": "Hello", "snippet": "Test email and stuff"  } ]`))
}

func handler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello World!\n"))
	return
}

// our main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/api/labels", getLabels)
	router.HandleFunc("/api/messages", getMessages)

	log.Fatal(http.ListenAndServe(":8000", router))
}
