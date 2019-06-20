package pinbox

import (
	notmuch "github.com/msp301/go.notmuch"
)

type Thread struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	NewestDate int64     `json:"newestDate"`
	OldestDate int64     `json:"oldestDate"`
	Authors    []string  `json:"authors"`
	Messages   []Message `json:"messages"`
}

func toOurThread(thr *notmuch.Thread) Thread {
	_, authors := thr.Authors()

	var messages []Message
	msg := notmuch.Message{}
	msgs := thr.Messages()
	for msgs.Next(&msg) {
		var files []string
		file := ""
		filenames := msg.Filenames()
		for filenames.Next(&file) {
			files = append(files, file)
		}

		message := Message{
			ID:    msg.ID(),
			Epoch: msg.Date().Unix(),
			Files: files,
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
	thr := notmuch.Thread{}
	for threads.Next(&thr) {
		res := toOurThread(&thr)
		payload = append(payload, res)
	}

	return payload
}
