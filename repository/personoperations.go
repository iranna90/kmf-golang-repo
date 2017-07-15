package repository

import (
	. "kmf-application/domain"
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
