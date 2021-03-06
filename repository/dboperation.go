package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=1111 dbname=kmfdetails user=kmfadmin  sslmode=disable password=changeme001")
	if err != nil {
		log.Panic("unable to create connection", err)
		panic(err)
	}
	return db
}
