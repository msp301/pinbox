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

type ourLabel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ourMessage struct {
	ID    string   `json:"id"`
	Epoch int64    `json:"epoch"`
	Files []string `json:"files"`
}

type ourThread struct {
	ID         string       `json:"id"`
	Subject    string       `json:"subject"`
	NewestDate int64        `json:"newestDate"`
	OldestDate int64        `json:"oldestDate"`
	Authors    []string     `json:"authors"`
	Messages   []ourMessage `json:messages`
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

	var payload []ourLabel
	for _, name := range tags.Slice() {
		label := ourLabel{ID: name, Name: name}
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

	query := db.NewQuery("*")
	threads, err := query.Threads()

	if err != nil {
		log.Println(err)
		return nil
	}

	var payload []ourThread
	thr := notmuch.Thread{}
	for threads.Next(&thr) {
		_, authors := thr.Authors()

		var messages []ourMessage
		msg := notmuch.Message{}
		msgs := thr.TopLevelMessages()
		for msgs.Next(&msg) {
			var files []string
			file := ""
			filenames := msg.Filenames()
			for filenames.Next(&file) {
				files = append(files, file)
			}

			message := ourMessage{
				ID:    msg.ID(),
				Epoch: msg.Date().Unix(),
				Files: files,
			}

			messages = append(messages, message)
		}

		res := ourThread{
			ID:         thr.ID(),
			Subject:    thr.Subject(),
			NewestDate: thr.NewestDate().Unix(),
			OldestDate: thr.OldestDate().Unix(),
			Authors:    authors,
			Messages:   messages,
		}
		payload = append(payload, res)
	}

	content, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return nil
	}

	return content
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
