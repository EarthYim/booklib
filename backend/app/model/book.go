package model

type Book struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Author string
	Genre  string
	Year   string
	Tag    string
}

type Tabler interface {
	TableName() string
}

func (Book) TableName() string {
	return "book"
}
