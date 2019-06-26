package pinbox

// A Bundle represents a group of Threads.
type Bundle struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Date    int64    `json:"date"`
	Threads []Thread `json:"threads"`
}
