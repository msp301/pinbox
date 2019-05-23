package pinbox

type Config struct {
	Maildir string
	Port    int
	Inbox   string
	Bundle  []string
	Hidden  []string
}
