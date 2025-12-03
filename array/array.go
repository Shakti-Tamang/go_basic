package array

import "fmt"

func DisplayArray() {
	fmt.Println("Array Elements:")

}

func PracArray() {

	// slice in its nto array
	var a []int

	a = []int{10, 20, 30, 40, 50}

	fmt.Println("Array Elements:", a[0])

	var arr [7]int
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500
	arr[5] = 600
	arr[6] = 700

	fmt.Println("Array Elements:", arr)

}
