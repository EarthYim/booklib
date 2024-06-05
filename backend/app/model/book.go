package model

type Book struct {
	ID    int    `gorm:"id,primary_key,auto_increment"`
	Title string `gorm:"title"`
	Autor string `gorm:"author"`
	Genre string `gorm:"genre"`
	Year  string `gorm:"year"`
	Tag   string `gorm:"tag"`
}
