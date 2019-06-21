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

	mailbox := pinbox.CreateNotmuch()
	mailbox.DbPath = config.Maildir
	mailbox.ExcludeLabels = config.Hidden
	mailbox.Bundle = config.Bundle
	mailbox.InboxLabel = config.Inbox

	api := pinbox.CreateMailboxAPI(mailbox)

	router := mux.NewRouter()
	router.UseEncodedPath()

	router.HandleFunc("/api/inbox", api.GetInbox)
	router.HandleFunc("/api/labels", api.GetLabels)
	router.HandleFunc("/api/messages/{id}", api.HandleSingleMessage)
	router.HandleFunc("/api/messages", api.HandleAllMessages)
	router.Path("/api/messages").Queries("label", "{label}").HandlerFunc(api.HandleLabeledMessages)

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, router))
}
