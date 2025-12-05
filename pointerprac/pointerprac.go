package pointerprac

import "fmt"

type Student struct {
	Name string

	Address string

	Roll string

	ContactNumber string
}

var (
	a *int

	b *int

	c []*int
)

func GetPointerValue() {

	// fmt.Println("pointer starting")

	n1 := 80

	n2 := 80

	a = &n1

	b = &n2

	c := *a + *b

	fmt.Println(c)

	fmt.Println("pointer")

}
