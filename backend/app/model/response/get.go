package response

type Get struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"Author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
	Tag    string `json:"tag"`
}
