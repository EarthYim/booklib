package request

type Edit struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Autor string `json:"autor"`
	Genre string `json:"genre"`
	Year  string `json:"year"`
	Tag   string `json:"tag"`
}
