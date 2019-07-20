package pinbox

// A Label is a tag that can be used to organise messages in the Mailbox.
type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
