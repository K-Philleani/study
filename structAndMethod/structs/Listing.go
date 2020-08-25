package structs

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}

func UpPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func Listing() {
	// 值类型 person
	var person1 Person
	person1.firstName = "Chris"
	person1.lastName = "Woodward"
	UpPerson(&person1)
	fmt.Printf("The name of the person is %s %s \n", person1.firstName, person1.lastName)

	// 指针类型 person
	person2 := new(Person)
	person2.firstName = "Chris"
	person2.lastName = "Woodward"
	(*person2).lastName = "Woodward!!!!"
	UpPerson(person2)
	fmt.Printf("The name of the person is %s %s\n", person2.firstName, person2.lastName)

	// literal person

	person3 := &Person{"Chris", "Woodward@@@"}
	fmt.Printf("The name of the person is %s %s \n", person3.firstName, person3.lastName)
}
