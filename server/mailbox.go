package pinbox

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"time"
)

// The Mailbox interface acts as a wrapper for all email actions to abstract any
// direct email access from the business logic of the Pinbox API transport.
type Mailbox interface {
	Labels() ([]Label, error)
	ReadMessage(id string) (MessageContent, error)
	Search(query string) ([]Thread, error)
}

// Inbox returns all email messages that match the configured 'Inbox' whilst
// grouping together any messages that match the labels configured in the 'Bundle' list.
//
// Returns a mixed list of Thread and Bundle objects.
// The list is returned in reverse chronological order.
func Inbox(mailbox Mailbox, config Config) ([]interface{}, error) {
	inboxLabel := config.Inbox

	queryString := "tag:" + inboxLabel
	for _, label := range config.Bundle {
		queryString += " and not tag:" + label
	}

	threads, err := mailbox.Search(queryString)
	if err != nil {
		return nil, errors.New("Search failed: " + err.Error())
	}

	bundles := make(map[string][]*Bundle)
	for _, label := range config.Bundle {
		res, err := mailbox.Search("tag:" + inboxLabel + " and tag:" + label)
		if err != nil {
			log.Println(fmt.Sprintf("Error with bundle '%s': %s", label, err))
			continue
		}

		bundled := make([]Thread, 0)
		var latestDate int64
		for _, thread := range res {
			date := thread.NewestDate

			if date > latestDate {
				latestDate = date
			}

			bundled = append(bundled, thread)
		}

		if len(bundled) > 0 {
			latest := time.Unix(latestDate, 0)
			month := fmt.Sprintf("%d %d", latest.Month(), latest.Year())
			bundles[month] = append(
				bundles[month],
				&Bundle{
					ID:      label,
					Type:    "bundle",
					Date:    latestDate,
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
	var prevDate time.Time
	for _, thread := range threads {
		date := time.Unix(thread.NewestDate, 0)
		month := fmt.Sprintf("%d %d", date.Month(), date.Year())

		if len(bundles[month]) > 0 {
			for _, bundle := range bundles[month] {
				bundleDate := bundle.Date

				if bundleDate > date.Unix() {
					bundle.Type = "bundle"
					inbox = append(inbox, bundle)
				}
			}
		}

		thread.Type = "thread"
		inbox = append(inbox, thread)
		prevDate = date
	}

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
				bundle.Type = "bundle"
				inbox = append(inbox, bundle)
			}
		}
	}

	return inbox, nil
}
