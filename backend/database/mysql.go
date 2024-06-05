package database

import (
	"booklib/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(config config.Config) gorm.DB {
	dsn := config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return *db
}
