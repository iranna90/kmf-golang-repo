package domain

import "time"

type Person struct {
	Id          int64 `gorm:"column:id"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (p Person) TableName() string {
	return "persons"
}

type Address struct {
	Id          int64 `gorm:"column:id"`
	PersonRef   int64 `gorm:"column:person_ref"`
	PhoneNumber int64 `gorm:"column:phone_number"`
	FullAddress string `gorm:"column:full_address"`
}

func (a Address) TableName() string {
	return "address"
}