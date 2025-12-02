package main

import "fmt"

// simple fucntion

func name(name string) string {
	fmt.Println("Name is:", name)
	return name
}

//2. Variadic Functions:
// These functions accept a variable number of arguments of a specific type. The variadic parameter is indicated by ... before the type in the function signature, and it's treated as a slice within the function body.

func displayVariadic(nums ...int) int {
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

// 3.// Anonymous Functions (Function Literals/Lambdas)
// These functions are defined without a name and can be assigned to variables, passed as arguments to other functions, or returned from functions. They are often used for short, one-time operations or as closures.
var anonFunc = func(msg string) {
	fmt.Println("Anonymous Function says:", msg)
}

// 4. Closure Functions:
// A closure is an anonymous function that "closes over" variables from its surrounding scope, even after the outer function has finished executing. This allows the inner function to retain access to those variables.

func counter() func() int {
	i := 0
	return func() int {
		i++

		return i
	}
}

// Higher-Order Functions:
// While not a distinct type of function in terms of declaration, Go supports higher-order functions, which are functions that take other functions as arguments or return functions as results. This is enabled by Go's support for function types and values.
