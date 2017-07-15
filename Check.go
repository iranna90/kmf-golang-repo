package main

import (
	"fmt"
	"kmf-application/repository"
	. "kmf-application/domain"
)

func main() {
	fmt.Println("hello")
	person := repository.RetrievePerson("Iranna", "Patil")
	fmt.Println("retrieved person ", person)
	address := Address{PhoneNumber: 9886881025, FullAddress: "Bangalore 566866", Person: person}
	repository.CreateAddress(&address)
	fmt.Println("stored ")
}
