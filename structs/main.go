package main

import "fmt"

func main() {

	p := Person{
		firstName: "bedoo",
		lastName:  "dev",
		age:       25,
		address:   Address{},
		Phone:     Phone{},
	}
	p.address.city = "Istanbul"
	p.address.country = "Turkey"
	p.home = "12345" // anonim embed edilen struct'lara doğrudan ulaşılabilir
	p.cell = "123456"

	fmt.Println(p)               // {bedoo dev 25 {Istanbul Turkey} {12345 123456}}
	fmt.Println(p.getFullName()) // bedoo dev

	// Anonim struct
	// user := struct {
	// 	username string
	// 	email    string
	// }{username: "bedoodev", email: "bedoodev@gmail.com"}
	// fmt.Println(user) // {bedoodev bedoodev@gmail.com}

	p.happyBirthDay()
	fmt.Println(p.age)
}

type Person struct {
	firstName string
	lastName  string
	age       uint
	address   Address
	Phone     // anonymous embedded
}

type Address struct {
	city    string
	country string
}

type Phone struct {
	home string
	cell string
}

func (p Person) getFullName() string {
	return fmt.Sprintln(p.firstName, p.lastName)
}

func (p *Person) happyBirthDay() {
	p.age++
}
