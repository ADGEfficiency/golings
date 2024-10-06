// structs2
// Make me compile!
package main

import "fmt"

type ContactDetails struct {
	phone string
}

type Person struct {
	details ContactDetails
	name    string
	age     int
}

func main() {
	contactDetails := ContactDetails{phone: "021"}
	person := Person{name: "John", age: 32, details: contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.details.phone)
}
