package pinbox

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/jhillyerd/enmime"
	notmuch "github.com/msp301/go.notmuch"
)

type Notmuch struct {
	DbPath        string
	ExcludeLabels []string
	InboxLabel    string
	Bundle        []string
}

func CreateNotmuch() *Notmuch {
	mailbox := Notmuch{}
	return &mailbox
}

func (mailbox *Notmuch) Inbox() ([]interface{}, error) {
	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	inboxLabel := mailbox.InboxLabel
	queryString := "tag:" + inboxLabel
	for _, label := range mailbox.Bundle {
		queryString += " and not tag:" + label
	}

	query := db.NewQuery(queryString)
	threads, err := query.Threads()
	defer threads.Close()

	bundles := make(map[string][]*Bundle)
	for _, label := range mailbox.Bundle {
		bundle := db.NewQuery("tag:" + inboxLabel + " and tag:" + label)
		res, err := bundle.Threads()

		if err != nil {
			log.Println(fmt.Sprintf("Error with bundle '%s': %s", label, err))
			continue
		}

		bundled := make([]*Thread, 0)
		var latestDate time.Time
		thread := notmuch.Thread{}
		for res.Next(&thread) {
			date := thread.NewestDate()

			if date.Unix() > latestDate.Unix() {
				latestDate = date
			}

			thr := toOurThread(&thread)
			bundled = append(bundled, &thr)
		}

		if len(bundled) > 0 {
			month := fmt.Sprintf("%d %d", latestDate.Month(), latestDate.Year())
			bundles[month] = append(
				bundles[month],
				&Bundle{
					ID:      label,
					Type:    "bundle",
					Date:    latestDate.Unix(),
					Threads: bundled,
				})
		}
	}

	for key := range bundles {
		sort.Slice(bundles[key], func(i, j int) bool {
			return bundles[key][i].Date > bundles[key][j].Date
		})
	}

	inbox := make([]interface{}, 0)
	thread := notmuch.Thread{}
	var prevDate time.Time
	for true {
		if !threads.Next(&thread) {
			// We have already added all top-level threads.
			// There may still be bundles that have not been included.
			// Flush out any remaining bundles that are older than our last top-level thread.
			keys := make([]string, 0)
			for key := range bundles {
				keys = append(keys, key)
			}
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))

			for _, key := range keys {
				for _, bundle := range bundles[key] {
					if bundle.Date < prevDate.Unix() {
						inbox = append(inbox, bundle)
					}
				}
			}

			break
		}

		date := thread.NewestDate()
		month := fmt.Sprintf("%d %d", date.Month(), date.Year())

		if len(bundles[month]) > 0 {
			for _, bundle := range bundles[month] {
				bundleDate := bundle.Date

				if bundleDate > date.Unix() {
					inbox = append(inbox, bundle)
				}
			}
		}

		thr := toOurThread(&thread)
		thr.Type = "thread"
		inbox = append(inbox, &thr)
		prevDate = date
	}

	return inbox, nil
}

func (mailbox *Notmuch) Labels() ([]Label, error) {
	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	tags, err := db.Tags()
	if err != nil {
		log.Println("Error getting tags")
	}

	hidden := make(map[string]int, 0)
	for _, label := range mailbox.ExcludeLabels {
		hidden[label] = 1
	}

	var payload []Label
	tag := notmuch.Tag{}
	for tags.Next(&tag) {
		name := tag.Value
		if hidden[name] == 1 {
			continue
		}
		label := Label{ID: name, Name: name}
		payload = append(payload, label)
	}

	return payload, nil
}

func (mailbox *Notmuch) ReadMessage(id string) (MessageContent, error) {

	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return MessageContent{}, err
	}

	defer db.Close()

	var msg *notmuch.Message
	var msgFilename string
	if err == nil {
		msg, err = db.FindMessage(id)

		if err == nil {
			msgFilename = msg.Filename()
		}
	}

	if err != nil {
		log.Println(fmt.Sprintf("%s: %s", id, err))
		return MessageContent{}, err
	}

	file, err := os.Open(msgFilename)

	if err != nil {
		log.Println(err)
		return MessageContent{}, err
	}
	defer file.Close()

	env, err := enmime.ReadEnvelope(file)

	if err != nil {
		log.Println(err)
		return MessageContent{}, err
	}

	body := env.HTML
	if len(body) == 0 {
		body = env.Text
	}

	encodedContent := base64.StdEncoding.EncodeToString([]byte(body))

	payload := MessageContent{
		ID:      msg.ID(),
		Epoch:   msg.Date().Unix(),
		Author:  msg.Header("From"),
		Content: encodedContent,
	}

	msg.Close()

	return payload, nil
}

func (mailbox *Notmuch) Search(query string) ([]Thread, error) {

	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	dbQuery := db.NewQuery(query)
	threads, err := dbQuery.Threads()

	payload := toOurThreads(threads)

	return payload, nil
}

func openDatabase(path string) (*notmuch.DB, error) {
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
		return nil, errors.New("Failed to open database")
	}

	return db, nil
}
