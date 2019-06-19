package pinbox

type Mailbox interface {
	Inbox() []interface{}
	Labels() []Label
	ReadMessage(id string) Message
	Search(query string) []Thread
}
