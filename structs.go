package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	Phone
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "New york",
			country: "USA",
		},
		Phone: Phone{
			number: "935793",
		},
	}
	p1 := Person{
		firstName: "JohnTerry",
		lastName:  "Doe",
	}
	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p1.fullName())
	p1.incrementAgeByOne()
	fmt.Println(p1.age)
	fmt.Println(p.address.city)
	fmt.Println(p.number)
	// Anonymous Structs
	// user := struct{
	// 	username string
	// 	email string
	// }{
	// 	username: "user123",
	// 	email:"pseudoemail@example.com",
	// }

}

func (a Person) fullName() string {
	return a.firstName + " " + a.lastName
}

func (a *Person) incrementAgeByOne() {
	a.age++
}

// struct embedding

type Address struct {
	city    string
	country string
}

type Phone struct {
	number string
}
