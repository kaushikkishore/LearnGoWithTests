package main

import "fmt"

// Create a struct name Person. Person have firstname and lastname with string property
type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (jimPtr *person) updateName(newName string) {
	(*jimPtr).firstname = newName
}

func main() {
	// Create values for struct
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
	}
	fmt.Println(alex)

	// 2nd option to define
	// Not a good choice.
	// john := person{"John", "Doe"}

	// 3rd way to define
	var john person

	john.firstname = "John"
	john.lastname = "Doe"

	fmt.Println(john)

	fmt.Printf("%+v", john)
	fmt.Println("")

	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contact: contactInfo{
			email:   "eWwZ7@example.com",
			zipcode: 12345,
		},
	}

	// one way
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// This will also work
	jim.updateName("jimmy")
	jim.print()

}
