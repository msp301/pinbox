package pinbox

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

type Mailbox interface {
	Inbox() ([]interface{}, error)
	Labels() ([]Label, error)
	ReadMessage(id string) (MessageContent, error)
	Search(query string) ([]Thread, error)
}

// MailboxAPI Contains logic for handling mailbox http requests
type MailboxAPI struct {
	mailbox Mailbox
}

func CreateMailboxAPI(mailbox Mailbox) *MailboxAPI {
	return &MailboxAPI{
		mailbox: mailbox,
	}
}

func (m *MailboxAPI) GetInbox(writer http.ResponseWriter, req *http.Request) {

	inbox, err := m.mailbox.Inbox()
	if err != nil {
		log.Println(err)
		return
	}

	content, err := json.Marshal(inbox)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		http.Error(writer, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	handler(content, writer)
}

func (m *MailboxAPI) GetLabels(writer http.ResponseWriter, req *http.Request) {

	payload, err := m.mailbox.Labels()
	if err != nil {
		log.Println("Failed to get labels", err)
		return
	}

	content, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		http.Error(writer, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	handler(content, writer)
}

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

	content, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		return
	}

	handler(content, writer)
}

func (m *MailboxAPI) HandleAllMessages(writer http.ResponseWriter, req *http.Request) {
	payload, err := m.mailbox.Search("*")
	if err != nil {
		log.Println("Failed to get messages", err)
		http.Error(writer, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	content, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		http.Error(writer, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	handler(content, writer)
}

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

	content, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to convert to JSON", err)
		http.Error(writer, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	handler(content, writer)
}

func handler(body []byte, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
	return
}
