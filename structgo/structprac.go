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

	first := Person{"shakti", 12, 13.4, "working as front end developer", false}

	fmt.Println(first)

}
