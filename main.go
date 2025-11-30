package main

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

func main() {
	logMessage("Application starting")
	requestHandler()
	logMessage("Application ending")
}

// Go needs a module (a go.mod file) because it needs to know which folder is your project and how to manage multiple files and packages.

// Without a module, Go doesn’t know which files belong together.

// That’s why it shows:

// Why You Should Use Go Modules

// Dependency Management: Gin and other libraries are installed via modules.

// Version Control: Modules allow you to specify exact versions of packages.

// Project Structure: Makes your project self-contained and reproducible.

// Required for Modern Go: Since Go 1.16+, modules are the standard way to manage dependencies.
