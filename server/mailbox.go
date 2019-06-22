package pinbox

// The Mailbox interface acts as a wrapper for all email actions to abstract any
// direct email access from the business logic of the Pinbox API transport.
type Mailbox interface {
	Inbox() ([]interface{}, error)
	Labels() ([]Label, error)
	ReadMessage(id string) (MessageContent, error)
	Search(query string) ([]Thread, error)
}
