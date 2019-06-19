package pinbox

import (
	"errors"
	"log"
	"os"

	notmuch "github.com/msp301/go.notmuch"
)

type Notmuch struct {
	DbPath        string
	ExcludeLabels []string
}

//func Inbox() []interface{} {}

func (mailbox *Notmuch) Labels() ([]Label, error) {
	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return nil, err
	}

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

//func ReadMessage(id string) Message {}

//func Search(query string) []Thread {}

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
