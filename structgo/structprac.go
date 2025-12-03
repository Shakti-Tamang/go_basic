package structgo

import "fmt"

// In Go, a struct (short for structure) is a user-defined type that groups together fields (variables) of potentially different data types into a single unit. It is a way to create complex data types that represent real-world entities.

// grouping related data: Structs allow you to collect related pieces of information into a single, cohesive unit. For example, instead of having separate name, age, and job variables, you can create a Person struct that encapsulates all of them.

// always must be starting with capital

// Go does not use public, private, or protected like Java or C#. Instead:

// Capital letter → means "exported" (like public)

// Lowercase → means "unexported" (like private)

// to make anything accvessible al over make it capital

type Person struct {
	Name          string
	Age           int16
	Salary        float64
	Job           string
	MaritalStatus bool
}

func StructSimple() {

	var r []Person

	var s Person

	fmt.Println(s)

	var p = Person{"mau", 12, 12.3, "wrking as designer", true}

	p.Name = "ram"

	fmt.Println(p.Name)

	addresspointer := new(Person)

	fmt.Println("pointer address", addresspointer)

	first := Person{"shakti", 12, 13.4, "working as front end developer", false}

	fmt.Println("age is" + first.Job)
	fmt.Println(first)

	last := Person{Name: "sita", Age: 12, Salary: 12.12, Job: "he  is figma designer", MaritalStatus: true}

	fmt.Println(last)

	laststruc := &last
	fmt.Println("string ", laststruc)

	r = append(r, p, first, last)

	for _per := range r {

		fmt.Println("all ages", r[_per].Age)

	}

	pp := Person{Name: "shakti", Job: "i amd dveveloper"}

	fmt.Println(pp.NameJob())

}

// Functions can be associated with a struct type, becoming "methods" of that struct. This allows structs to have behavior in addition to data.

func (p Person) NameJob() string {

	return p.Name + p.Job
}
