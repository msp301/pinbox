package pinbox

type Mailbox interface {
	Inbox() ([]interface{}, error)
	Labels() ([]Label, error)
	ReadMessage(id string) Message
	Search(query string) ([]Thread, error)
}
