package main

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
