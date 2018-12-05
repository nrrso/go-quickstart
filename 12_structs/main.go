package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	//person1 := Person{firstName: "Norris", lastName: "Osarenkhoe", city: "Hamburg", gender: "m", age: 27}

	// alternative only useful when loading via vars
	person2 := Person{"Norris", "Osarenkhoe", "Hamburg", "m", 27}
	person1 := Person{"Samantha", "Williams", "Hamburg", "f", 27}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Mueller")
	fmt.Println(person2.greet())
}
