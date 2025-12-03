package structgo

import "fmt"

// In Go, a struct (short for structure) is a user-defined type that groups together fields (variables) of potentially different data types into a single unit. It is a way to create complex data types that represent real-world entities.

// grouping related data: Structs allow you to collect related pieces of information into a single, cohesive unit. For example, instead of having separate name, age, and job variables, you can create a Person struct that encapsulates all of them.
type Person struct {
	name string

	age int16

	salary float64

	job string

	maritialstatus bool
}

func StructSimple() {

	var r []Person

	var p Person

	p = Person{"mau", 12, 12.3, "wrking as designer", true}

	p.name = "ram"

	fmt.Println(p.name)

	addresspointer := new(Person)

	fmt.Println("pointer address", addresspointer)

	first := Person{"shakti", 12, 13.4, "working as front end developer", false}

	fmt.Println("age is" + first.job)
	fmt.Println(first)

	last := Person{name: "sita", age: 12, salary: 12.12, job: "he  is figma designer", maritialstatus: true}

	fmt.Println(last)

	laststruc := &last
	fmt.Println("string ", laststruc)

	r = append(r, p, first, last)

	for _per := range r {

		fmt.Println("all ages", r[_per].age)

	}

}

// Functions can be associated with a struct type, becoming "methods" of that struct. This allows structs to have behavior in addition to data.
