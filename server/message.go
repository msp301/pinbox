package pinbox

// A Message represents an individual email.
type Message struct {
	ID     string   `json:"id"`
	Epoch  int64    `json:"epoch"`
	Author string   `json:"author"`
	Files  []string `json:"files"`
}

// MessageContent is a container for the body of a Message.
type MessageContent struct {
	ID      string `json:"id"`
	Epoch   int64  `json:"epoch"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
