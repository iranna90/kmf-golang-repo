package repository

import (
	. "kmf-application/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func RetrievePerson(firstName, lastName string) (person Person) {
	db := getDatabaseConnection()
	defer db.Close()

	db.Where(&Person{FirstName: firstName, LastName: lastName}).First(&person)
	return
}

func CreatePerson(person *Person) {
	db := getDatabaseConnection()
	defer db.Close()

	db.Create(person)
}

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=1111 dbname=kmfdetails user=kmfadmin  sslmode=disable password=changeme001")
	if err != nil {
		log.Panic("unable to create connection", err)
		panic(err)
	}
	return db
}
