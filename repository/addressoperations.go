package repository

import (
	. "kmf-application/domain"
)

func RetrieveAddress(personRef int64) (address Address) {
	db := getDatabaseConnection()
	defer db.Close()

	db.Where(&Address{PersonRef: personRef}).First(&address)
	return
}

func CreateAddress(address *Address) {
	db := getDatabaseConnection()
	defer db.Close()

	db.Create(address)
}
