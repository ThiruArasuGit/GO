package main

import (
	"fmt"
	"strconv"
)

//Its like class concepts- Actually there are no class in GO

//define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//Method: Value receiver
func (p Person) greet() string {
	return "Hello!, My name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

//Method: pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

//Method; getMarried (Pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	//init Person using struct
	person1 := Person{firstName: "Thiru", lastName: "A", city: "Bangalore", gender: "M", age: 30}
	fmt.Println(person1)

	person2 := Person{"Ashwini", "P", "NewYork", "F", 28}
	fmt.Println(person2)

	person1.getMarried("Ashwini")
	person2.getMarried("Thiru")

	fmt.Println(person1.firstName + ", " + person2.firstName)
	//Value receiver: = akes the value and does not make any changes.
	fmt.Println(person2.greet())
	fmt.Println(person1.greet())

	//Pointer receiver := Will make some changes.
	person1.hasBirthday()
	person2.hasBirthday()

	fmt.Println(person1.age)
	fmt.Println(person2.age)

}
