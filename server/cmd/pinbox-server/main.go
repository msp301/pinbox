package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"

	"github.com/msp301/pinbox-server"
)

func getLabels(writer http.ResponseWriter, req *http.Request, mailbox pinbox.Mailbox) {

	payload, err := mailbox.Labels()

	if err != nil {
		log.Println("Failed to get labels", err)
		return
	}

	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return
	}

	handler(content, writer)
}

func getMessage(writer http.ResponseWriter, req *http.Request, mailbox pinbox.Mailbox) {
	vars := mux.Vars(req)
	rawID, err := url.PathUnescape(vars["id"])

	if err != nil {
		log.Println("Failed to extract message ID from URL", err)
		return
	}

	payload, err := mailbox.ReadMessage(rawID)

	if err != nil {
		log.Println("Failed to read message", err)
		return
	}

	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return
	}

	handler(content, writer)
}

func getMessages(writer http.ResponseWriter, req *http.Request, mailbox pinbox.Mailbox, query string) {
	payload, err := mailbox.Search(query)

	if err != nil {
		log.Println("Failed to get messages", err)
		return
	}

	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return
	}

	handler(content, writer)
}

func handler(body []byte, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
	return
}

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

	router := mux.NewRouter()

	router.UseEncodedPath()

	router.HandleFunc("/api/inbox", func(writer http.ResponseWriter, req *http.Request) {

		inbox, err := mailbox.Inbox()

		if err != nil {
			log.Println(err)
			return
		}

		content, err := json.Marshal(inbox)
		if err != nil {
			log.Println("Failed to convert to JSON", err)
			return
		}

		handler(content, writer)
	})

	router.HandleFunc("/api/labels", func(writer http.ResponseWriter, req *http.Request) {
		getLabels(writer, req, mailbox)
	})

	router.Path("/api/messages").Queries("label", "{label}").HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		url := req.URL
		labels := url.Query()["label"]

		var dbQuery = ""
		for _, label := range labels {
			if len(label) == 0 {
				continue
			}

			if len(dbQuery) == 0 {
				dbQuery = dbQuery + " tag:" + label
			} else {
				dbQuery = dbQuery + " and tag:" + label
			}
		}

		if len(dbQuery) == 0 {
			log.Println("No labels given")
			return
		}

		if err != nil {
			log.Println(err)
			return
		}

		getMessages(writer, req, mailbox, dbQuery)
	})

	router.HandleFunc("/api/messages", func(writer http.ResponseWriter, req *http.Request) {
		getMessages(writer, req, mailbox, "*")
	})

	router.HandleFunc("/api/messages/{id}", func(writer http.ResponseWriter, req *http.Request) {
		getMessage(writer, req, mailbox)
	})

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, router))
}
