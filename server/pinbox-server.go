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
	"time"

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

type ourBundle struct {
	ID      string       `json:"id"`
	Threads []*ourThread `json:"threads"`
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
	tag := notmuch.Tag{}
	for tags.Next(&tag) {
		name := tag.Value
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

func toOurThread(thr *notmuch.Thread) ourThread {
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

	return res
}

func toOurThreads(threads *notmuch.Threads) []ourThread {
	var payload []ourThread
	thr := notmuch.Thread{}
	for threads.Next(&thr) {
		res := toOurThread(&thr)
		payload = append(payload, res)
	}

	return payload
}

func getMessages(writer http.ResponseWriter, req *http.Request, db *notmuch.DB, threads *notmuch.Threads) {
	payload := toOurThreads(threads)
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

	router.HandleFunc("/api/inbox", func(writer http.ResponseWriter, req *http.Request) {
		query := db.NewQuery("tag:inbox and not tag:attachment")
		threads, err := query.Threads()

		labels := []string{"attachment"}
		bundles := make(map[string]*ourBundle)
		for _, label := range labels {
			bundle := db.NewQuery("tag:" + label)
			res, err := bundle.Threads()

			if err != nil {
				log.Println(fmt.Sprintf("Error with bundle '%s': %s", label, err))
				continue
			}

			bundled := make([]*ourThread, 0)
			var latestDate time.Time
			thread := notmuch.Thread{}
			for res.Next(&thread) {
				date := thread.NewestDate()

				if date.Year() > latestDate.Year() && date.Month() > latestDate.Month() {
					latestDate = date
				}

				thr := toOurThread(&thread)
				bundled = append(bundled, &thr)
			}

			month := fmt.Sprintf("%d %d", latestDate.Month(), latestDate.Year())
			bundles[month] = &ourBundle{ID: label, Threads: bundled}
		}

		inbox := make([]interface{}, 0)
		prevMonth := ""
		thread := notmuch.Thread{}
		for threads.Next(&thread) {
			date := thread.NewestDate()
			month := fmt.Sprintf("%d %d", date.Month(), date.Year())

			if month != prevMonth {
				if bundles[prevMonth] != nil {
					inbox = append(inbox, bundles[prevMonth])
				}
			}

			thr := toOurThread(&thread)
			inbox = append(inbox, &thr)
			prevMonth = month
		}

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
		threads.Close()
	})

	router.HandleFunc("/api/labels", func(writer http.ResponseWriter, req *http.Request) {
		getLabels(writer, req, db)
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

		query := db.NewQuery(dbQuery)
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
