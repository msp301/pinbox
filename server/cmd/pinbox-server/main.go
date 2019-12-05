package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/msp301/pinbox-server"
)

func usage() string {
	return fmt.Sprintf(`Usage: %s <config_file>`, os.Args[0])
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		println(usage())
		os.Exit(0)
	}

	configPath := args[0]

	config, err := pinbox.ReadConfigFile(configPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
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

	if config.TLS {
		err = http.ListenAndServeTLS(port, config.CertificateFile, config.CertificateKey, router)
	} else {
		err = http.ListenAndServe(port, router)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
