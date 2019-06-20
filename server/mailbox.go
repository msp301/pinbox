package pinbox

type Mailbox interface {
	Inbox() []interface{}
	Labels() ([]Label, error)
	ReadMessage(id string) Message
	Search(query string) []Thread
}
