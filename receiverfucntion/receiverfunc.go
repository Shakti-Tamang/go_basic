package receiverfucntion

import (
	"fmt"
)

type UserDetails struct {
	Name   string
	Age    int
	Salary float32
}

func (use *UserDetails) UserCustom() {

	fmt.Println("user details ", use.Age, use.Name)
}

func GetUsers(users *UserDetails) {
	p := UserDetails{Name: "shakti", Age: 23}

	p.UserCustom()

	if users.Age == 12 {
		fmt.Println("users", users.Age, users.Name, users.Salary)
	} else {
		fmt.Println("error")
	}

	// Check passed user
	if users.Age == 12 {
		fmt.Println("Matched User:", users.Name, users.Age, users.Salary)
	} else {
		fmt.Println("Error: Age did not match")
	}

}
