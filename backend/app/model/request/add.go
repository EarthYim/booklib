package request

type Add struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
	Tag    string `json:"tag"`
}
