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
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/gorilla/mux"
)

// MailboxAPI contains logic for handling mailbox http requests
type MailboxAPI struct {
	mailbox Mailbox
	config  Config
}

// CreateMailboxAPI creates a new Mailbox API instance.
// Return a new MailboxAPI reference.
func CreateMailboxAPI(mailbox Mailbox, config Config) *MailboxAPI {
	return &MailboxAPI{
		mailbox: mailbox,
		config:  config,
	}
}

// GetInbox retrieves all inbox messages in the Mailbox.
func (m *MailboxAPI) GetInbox(writer http.ResponseWriter, req *http.Request) {

	inbox, err := Inbox(m.mailbox, m.config)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error occurred accessing inbox", http.StatusInternalServerError)
		return
	}

	listHandler(inbox, writer)
}

// GetLabels retrieves the available labels in the Mailbox.
func (m *MailboxAPI) GetLabels(writer http.ResponseWriter, req *http.Request) {

	payload, err := m.mailbox.Labels()
	if err != nil {
		log.Println("Failed to get labels", err)
		http.Error(writer, "Failed to get labels", http.StatusInternalServerError)
		return
	}

	listHandler(payload, writer)
}

// HandleSingleMessage retrieves a message by ID from the Mailbox.
func (m *MailboxAPI) HandleSingleMessage(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	rawID, err := url.PathUnescape(vars["id"])

	if err != nil {
		log.Println("Failed to extract message ID from URL", err)
		http.Error(writer, "Failed to extract message ID from URL", http.StatusBadRequest)
		return
	}

	payload, err := m.mailbox.ReadMessage(rawID)
	if err != nil {
		log.Println("Failed to read message", err)
		http.Error(writer, "Failed to read message", http.StatusInternalServerError)
		return
	}

	handler(payload, writer)
}

// HandleAllMessages retrieves all messages in the Mailbox.
func (m *MailboxAPI) HandleAllMessages(writer http.ResponseWriter, req *http.Request) {
	payload, err := m.mailbox.Search("*")
	if err != nil {
		log.Println("Failed to get messages", err)
		http.Error(writer, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	listHandler(payload, writer)
}

// HandleLabeledMessages retrieves any messages in the Mailbox with the specified labels.
func (m *MailboxAPI) HandleLabeledMessages(writer http.ResponseWriter, req *http.Request) {
	labels := req.URL.Query()["label"]

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
		http.Error(writer, "No labels given", http.StatusBadRequest)
		return
	}

	payload, err := m.mailbox.Search(dbQuery)
	if err != nil {
		log.Println("Failed to get messages", err)
		http.Error(writer, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	listHandler(payload, writer)
}

func handler(content interface{}, writer http.ResponseWriter) {
	body, err := json.Marshal(content)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		http.Error(writer, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
	return
}

func listHandler(content interface{}, writer http.ResponseWriter) {
	// Generally we're expecting a series of results.
	// To stop json.Marshal outputting 'null', assume the intent is to return a list.
	if reflect.ValueOf(content).IsNil() {
		content = make([]interface{}, 0)
	}

	handler(content, writer)
	return
}
