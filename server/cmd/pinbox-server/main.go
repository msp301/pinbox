package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/msp301/pinbox-server"
)

func main() {
	args := os.Args[1:]
	configPath := args[0]

	config, err := pinbox.ReadConfigFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mailbox := pinbox.CreateNotmuch(config)

	api := pinbox.CreateMailboxAPI(mailbox, config)

	router := mux.NewRouter()
	router.UseEncodedPath()

	router.Path("/api/messages").Queries("label", "{label}").HandlerFunc(api.HandleLabeledMessages)

	router.HandleFunc("/api/inbox", api.GetInbox)
	router.HandleFunc("/api/labels", api.GetLabels)
	router.HandleFunc("/api/messages/{id}", api.HandleSingleMessage)
	router.HandleFunc("/api/messages", api.HandleAllMessages)

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, router))
}
