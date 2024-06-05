package request

type Add struct {
	Title string `json:"title"`
	Autor string `json:"autor"`
	Genre string `json:"genre"`
	Year  string `json:"year"`
	Tag   string `json:"tag"`
}
