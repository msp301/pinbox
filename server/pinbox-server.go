package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	notmuch "github.com/msp301/go.notmuch"
)

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

func getLabels(db *notmuch.DB) []byte {
	tags, err := db.Tags()
	if err != nil {
		log.Println("Error getting tags")
	}

	var payload []Label
	for _, name := range tags.Slice() {
		label := Label{ID: name, Name: name}
		payload = append(payload, label)
	}
	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return nil
	}

	return content
}

func getMessages(db *notmuch.DB) []byte {
	return []byte(`[ { "id": 999, "epoch": 1550930545657, "recipient": "Test <hello@test.com>", "sender": "API <api@server.net>", "subject": "Hello", "snippet": "Test email and stuff"  } ]`)
}

func handler(body []byte) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Write(body)
		return
	}
}

// our main function
func main() {
	args := os.Args[1:]
	dir, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal(err)
	}

	db := openIndexDatabase(dir)
	router := mux.NewRouter()

	router.HandleFunc("/api/labels", handler(getLabels(db)))
	router.HandleFunc("/api/messages", handler(getMessages(db)))

	log.Fatal(http.ListenAndServe(":8000", router))
}
