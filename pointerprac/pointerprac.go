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

	ptr *[]int
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

func GetSlicePointer() {
	a := []int{1, 2, 3}

	ptr = &a

	// chnaging pointer slice value
	(*ptr)[0] = 100

	fmt.Println(*ptr)

}

func DoubleSlice(s *[]int) {
	*s = append(*s, *s...)

}
