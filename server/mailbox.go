package pinbox

type Mailbox interface {
	Inbox() ([]interface{}, error)
	Labels() ([]Label, error)
	ReadMessage(id string) (MessageContent, error)
	Search(query string) ([]Thread, error)
}
