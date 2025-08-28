package main

import "fmt"

func ex3() {
	// non-anonymous struct
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	bob := Employee{
		firstName: "Bob",
		lastName:  "Julie",
		id:        1,
	}

	// anonymous struct
	var person struct {
		firstName string
		lastName  string
		id        int
	}

	person.firstName = "Bob"
	person.lastName = "Julie"
	person.id = 2

	// literal
	john := struct {
		firstName string
		lastName  string
		id        int
	}{
		firstName: "John",
		lastName:  "wick",
		id:        3,
	}

	fmt.Println(bob)
	fmt.Println(person)
	fmt.Println(john)

	return
}
