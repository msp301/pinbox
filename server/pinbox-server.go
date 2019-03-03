package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jhillyerd/enmime"
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

type ourMessageContent struct {
	ID      string `json:"id"`
	Epoch   int64  `json:"epoch"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type ourThread struct {
	ID         string       `json:"id"`
	Subject    string       `json:"subject"`
	NewestDate int64        `json:"newestDate"`
	OldestDate int64        `json:"oldestDate"`
	Authors    []string     `json:"authors"`
	Messages   []ourMessage `json:"messages"`
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

func getLabels(writer http.ResponseWriter, req *http.Request, db *notmuch.DB) {
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
		msgFilename = msg.Filename()
	}

	if err != nil {
		log.Println(err)
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

	payload := ourMessageContent{
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

func getMessages(writer http.ResponseWriter, req *http.Request, db *notmuch.DB, threads *notmuch.Threads) {
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
		msgs.Close()

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
		return
	}

	handler(content, writer)
}

func handler(body []byte, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
	return
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

	router.UseEncodedPath()

	router.HandleFunc("/api/labels", func(writer http.ResponseWriter, req *http.Request) {
		getLabels(writer, req, db)
	})

	router.Path("/api/messages").Queries("label", "{label}").HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		query := db.NewQuery("tag:" + vars["label"])
		threads, err := query.Threads()

		if err != nil {
			log.Println(err)
			return
		}

		getMessages(writer, req, db, threads)
		threads.Close()
	})

	router.HandleFunc("/api/messages", func(writer http.ResponseWriter, req *http.Request) {
		query := db.NewQuery("*")
		threads, err := query.Threads()

		if err != nil {
			log.Println(err)
			return
		}

		getMessages(writer, req, db, threads)
		threads.Close()
	})

	router.HandleFunc("/api/messages/{id}", func(writer http.ResponseWriter, req *http.Request) {
		getMessage(writer, req, db)
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
