package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jhillyerd/enmime"
	notmuch "github.com/msp301/go.notmuch"

	"github.com/msp301/pinbox-server"
)

func openIndexDatabase(path string) *notmuch.DB {
	var db *notmuch.DB
	var status error
	var dbPath = path + "/.notmuch"

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		db, status = notmuch.Create(path)
		db.Close()
	}
	if status == nil {
		db, status = notmuch.Open(path, notmuch.DBReadWrite)
	} else {
		log.Fatal("Failed to open database")
	}

	return db
}

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

func getMessage(writer http.ResponseWriter, req *http.Request, db *notmuch.DB) {
	vars := mux.Vars(req)
	rawID, err := url.PathUnescape(vars["id"])

	var msg *notmuch.Message
	var msgFilename string
	if err == nil {
		msg, err = db.FindMessage(rawID)

		if err == nil {
			msgFilename = msg.Filename()
		}
	}

	if err != nil {
		log.Println(fmt.Sprintf("%s: %s", rawID, err))
		return
	}

	file, err := os.Open(msgFilename)

	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	env, err := enmime.ReadEnvelope(file)

	if err != nil {
		log.Println(err)
		return
	}

	body := env.HTML
	if len(body) == 0 {
		body = env.Text
	}

	encodedContent := base64.StdEncoding.EncodeToString([]byte(body))

	payload := pinbox.MessageContent{
		ID:      msg.ID(),
		Epoch:   msg.Date().Unix(),
		Author:  msg.Header("From"),
		Content: encodedContent,
	}

	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return
	}

	handler(content, writer)
	msg.Close()
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

	dir, err := filepath.Abs(config.Maildir)
	if err != nil {
		log.Fatal(err)
	}

	mailbox := pinbox.CreateNotmuch()
	mailbox.DbPath = config.Maildir
	mailbox.ExcludeLabels = config.Hidden
	mailbox.Bundle = config.Bundle
	mailbox.InboxLabel = config.Inbox

	db := openIndexDatabase(dir)
	db.Close()
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
		getMessage(writer, req, db)
	})

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, router))
}
