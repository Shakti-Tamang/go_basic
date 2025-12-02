package main

import "fmt"

// simple fucntion

func name(name string) string {
	fmt.Println("Name is:", name)
	return name
}

//2. Variadic Functions:
// These functions accept a variable number of arguments of a specific type. The variadic parameter is indicated by ... before the type in the function signature, and it's treated as a slice within the function body.

func displayVariadicI(nums ...int) int {
	var sum int

	// for _, num := range nums {

	// This is a for-each loop in Go.

	// Let’s break it:

	// for

	// → starting a loop.

	// _, num

	// Go gives 2 values when looping a slice:

	// index (0,1,2…)

	// value at that index

	// Here:

	// _ = blank → you say I don’t need the index

	// num = each number in the slice

	// :=

	// Short variable declaration.

	// range nums

	for _, num := range nums {
		sum += num
	}
	return sum

}
