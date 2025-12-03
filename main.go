package main

import (
	"fmt"
	"goprac/array"
	"strings"
)

// Go is a statically typed, compiled, and concurrent programming language designed by Google. It emphasizes simplicity, performance, and efficient handling of modern systems programming tasks.

// THERE IS GO COMPILER TO CIMPLE GO TO BINARY AND RUN IT

// go run hello.go → This command compiles and runs only the hello.go file. Go creates a temporary binary from hello.go and executes it immediately. After execution, the binary is deleted.

// go run . → This command compiles and runs all Go files in the current directory that belong to the main package. It’s used when your program has multiple .go files working together.

// COMMANDS:

// go mod init goprac

// go run .

// Since Go 1.11, Go uses modules to manage projects and dependencies.

// A module is a collection of Go packages in a directory with a go.mod file at its root.

// It tells Go the module name and dependencies.

// It allows you to run multiple files together and import packages reliably.

// Without a go.mod, Go doesn’t know:

// What is the root of your project

// What packages belong together

// Where to find dependencies

const (
	a          = 10
	role       = "admin"
	isverified = true
	salary     = 45000.50
)

var (
	b     int
	c     string
	d     bool
	e     float64
	f     int16
	uage  int
	uname string

	// FG  :=70  cannot work in package level
)
var increamets = 9

func main() {

	increamets = 50

	var shakti = "name"

	fmt.Println("shakti:", shakti)

	logMessage("Application starting")
	requestHandler()
	logMessage("Application ending")
	anu := strings.ToUpper("shakti")
	fmt.Println("Name:", anu)

	// it is variable

	// short variable declaration
	at := "amu"
	at = "saman"
	fmt.Println("at:", at)
	var b = sum(1, 2)
	fmt.Println("Sum:", b)

	details()

	displayDatas()

	// 	"Type inferred as int at compile time" means:

	// Type inferred

	// You did not explicitly write the type, or even if you did, Go knows the type from the value you assign.

	// Go automatically figures out the type from the value.

	// int

	// Go determined that the type of the value 60 is an integer (int).

	// At compile time

	// This happens while compiling the code, not at runtime.

	const roll int = 60 // Type inferred as int at compile time

	fmt.Println("Value of as:", roll)

	// In Go, the := syntax is called short variable declaration

	// 	It’s called a “short variable declaration” because it combines two things in a single, short syntax:

	// Declaring a new variable

	// Assigning it a value

	// 	This declares a new variable and assigns a value at the same time.

	// Go automatically infers the type from the value.

	// Works only inside functions, not at the package level.

	// Go infers the type automatically, so you don’t need to write int, string, etc.

	// it is only for variable declaration and initialization.

	// You can’t use := for constants; use const instead.
	name := "shakti"

	userage := 25

	fmt.Println("Name:", name)
	fmt.Println("Age:", userage)

	increment()

	fmt.Println("enter the name")

	// 	&name means “the memory address of the variable name”.

	// In Go, this is called a pointer.

	// 	A pointer is a variable that stores the address of another variable in memory.

	// Instead of storing the actual value, it stores where the value is located.

	fmt.Scan(&uname) // &name → pass pointer
	fmt.Scan()
	fmt.Println("enter the age")

	fmt.Scan(&uage)
	bool := getBooleanValue(uname)

	fmt.Println("Boolean Value:", bool, "\t", uage)

	var total = displayVariadic(1, 2, 3, 4, 5)
	fmt.Println("Total Sum:", total)

	anonFunc("Hello from Anonymous Function")

	c := counter()

	fmt.Println("Counter Value:", c())

	fmt.Println("Counter Value:", c())

	array.DisplayArray()

	array.PracArray()
}

func details() {

	fmt.Println("Role:", role)

	fmt.Println("Is Verified:", isverified)
	fmt.Println("Salary:", salary)
}

func sum(a int, b int) int {
	return a + b
}

func getBooleanValue(name string) bool {

	if name == "shakti" {
		return true
	}

	// return strings.EqualFold(name, "shakti")

	return false
}
func displayDatas() {

	b = 345
	c = "Shakti"
	d = true
	e = 45.67
	f = 1234
	fmt.Println("Integer:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
	fmt.Println("Float64:", e)
	fmt.Println("Int16:", f)
}

func increment() {
	increamets++
	fmt.Println("Incremented Value:", increamets)
}

// Go needs a module (a go.mod file) because it needs to know which folder is your project and how to manage multiple files and packages.

// Without a module, Go doesn’t know which files belong together.

// That’s why it shows:

// Why You Should Use Go Modules

// Dependency Management: Gin and other libraries are installed via modules.

// Version Control: Modules allow you to specify exact versions of packages.

// Project Structure: Makes your project self-contained and reproducible.

// Required for Modern Go: Since Go 1.16+, modules are the standard way to manage dependencies.
