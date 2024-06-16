package request

type Add struct {
	Title  string `json:"title"`
	Author string `json:"Author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
	Tag    string `json:"tag"`
}
