package domain

import "time"

type Person struct {
	Id          int64 `gorm:"column:id"`
	PersonId    string `gorm:"column:person_id"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	LastUpdated time.Time `gorm:"column:last_updated"`
}

func (p Person) TableName() string {
	return "persons"
}

type Address struct {
	Id          int64 `gorm:"column:id"`
	PhoneNumber int64 `gorm:"column:phone_number"`
	FullAddress string `gorm:"column:full_address"`
	Person      Person `gorm:"ForeignKey:PersonRef;AssociationForeignKey:Id"`
	PersonRef   int64 `gorm:"column:person_ref"`
}

func (a Address) TableName() string {
	return "address"
}

type Balance struct {
	Id        int64 `gorm:"column:id"`
	Amount    int64 `gorm:"column:amount"`
	Modified  time.Time `gorm:"column:last_updated"`
	Person    Person `gorm:"ForeignKey:PersonRef;AssociationForeignKey:Id"`
	PersonRef int64 `gorm:"column:person_ref"`
}

func (b Balance) TableName() string {
	return "total_balance"
}

type MilkTransaction struct {
	Id         int64 `gorm:"column:id"`
	Liters     int64 `gorm:"column:number_of_liters"`
	TotalPrice int64 `gorm:"column:total_price_of_day"`
	Day        time.Time `gorm:"column:day"`
	PersonName string `gorm:"column:person_name"`
	Person     Person `gorm:"ForeignKey:PersonRef;AssociationForeignKey:Id"`
	PersonRef  int64 `gorm:"column:person_ref"`
}

func (m MilkTransaction) TableName() string {
	return "daily_transactions"
}

type Payment struct {
	Id        int64 `gorm:"column:id"`
	amount    int64 `gorm:"column:amount_payed"`
	PaidTo    string `gorm:"column:paid_to"`
	Day       time.Time `gorm:"column:day"`
	Person    Person `gorm:"ForeignKey:PersonRef;AssociationForeignKey:Id"`
	PersonRef int64 `gorm:"column:person_ref"`
}

func (m Payment) TableName() string {
	return "payment_details"
}
