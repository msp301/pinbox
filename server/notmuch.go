// Copyright 2019-2022 Martin Pritchard
//
// This file is part of Pinbox.
//
// Pinbox is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
//
// Pinbox is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with Pinbox.  If not, see <https://www.gnu.org/licenses/>.

package pinbox

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhillyerd/enmime"
	notmuch "github.com/msp301/go.notmuch"
)

// Notmuch is a Mailbox object intended to interact with a Notmuch database.
// See https://notmuchmail.org/ for more information.
type Notmuch struct {
	DbPath        string
	ExcludeLabels []string
	InboxLabel    string
	Bundle        []string
}

// CreateNotmuch creates a new Notmuch object ready to interact with the
// database specified by the given configuration.
func CreateNotmuch(config Config) *Notmuch {
	mailbox := Notmuch{
		DbPath:        config.Maildir,
		ExcludeLabels: config.Hidden,
		InboxLabel:    config.Inbox,
		Bundle:        config.Bundle,
	}
	return &mailbox
}

// Labels returns a list of Label objects corresponding to 'tags' configured in
// the Notmuch database.
// Any label IDs specified in Notmuch.ExcludeLabels will be omitted.
func (mailbox *Notmuch) Labels() ([]Label, error) {
	db, err := openDatabase(mailbox.DbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	tags, err := db.Tags()
	if err != nil {
		return nil, errors.New("Error getting tags")
	}

	hidden := make(map[string]int, 0)
	for _, label := range mailbox.ExcludeLabels {
		hidden[label] = 1
	}

	var payload []Label
	tag := &notmuch.Tag{}
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

// ReadMessage retrieves the content of an email.
// Returns a MessageContent object containing email body encoded in base64.
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
		return MessageContent{}, fmt.Errorf("%s: %s", id, err)
	}

	file, err := os.Open(msgFilename)

	if err != nil {
		return MessageContent{}, err
	}
	defer file.Close()

	env, err := enmime.ReadEnvelope(file)

	if err != nil {
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

// Search retrieves emails from the Notmuch database based on a given query.
// The query string must be provided in Xapian query format (https://notmuchmail.org/searching/).
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

	_, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Mailbox directory '%s' does not exist", path)
	}
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

func toOurThread(thr *notmuch.Thread) Thread {
	_, authors := thr.Authors()

	var messages []Message
	msg := &notmuch.Message{}
	msgs := thr.Messages()
	for msgs.Next(&msg) {
		var files []string
		file := ""
		filenames := msg.Filenames()
		for filenames.Next(&file) {
			files = append(files, file)
		}

		message := Message{
			ID:     msg.ID(),
			Author: msg.Header("From"),
			Epoch:  msg.Date().Unix(),
			Files:  files,
		}

		messages = append(messages, message)
	}
	msgs.Close()

	res := Thread{
		ID:         thr.ID(),
		Subject:    thr.Subject(),
		NewestDate: thr.NewestDate().Unix(),
		OldestDate: thr.OldestDate().Unix(),
		Authors:    authors,
		Messages:   messages,
	}

	return res
}

func toOurThreads(threads *notmuch.Threads) []Thread {
	var payload []Thread
	thr := &notmuch.Thread{}
	for threads.Next(&thr) {
		res := toOurThread(thr)
		payload = append(payload, res)
	}

	return payload
}
