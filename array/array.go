package array

import "fmt"

func DisplayArray() {
	fmt.Println("Array Elements:")

}

func PracArray() {

	// slice in its nto array
	var a []int

	// size zero we cannot keep
	// This declares a slice, not an array.

	// Right now, it’s a nil slice — no memory is allocated yet for elements.
	a = []int{10, 20, 30, 40, 50}

	a = append(a, 2)

	fmt.Println("Array Elements:", a)

	var arr [7]int
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500
	arr[5] = 600
	arr[6] = 700

	for _num := range arr {

		fmt.Println("Array Element at index", _num, "is", arr[_num])
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Array Element at index", i, "is", arr[i])

	}

	fmt.Println("Array Elements:", arr)

	// inferred length array:

	arr2 := [...]string{"Go", "Python", "Java", "C++", "JavaScript"}

	fmt.Println("Inferred Length Array Elements:", arr2)

}
