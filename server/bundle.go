package pinbox

type Bundle struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Date    int64     `json:"date"`
	Threads []*Thread `json:"threads"`
}
