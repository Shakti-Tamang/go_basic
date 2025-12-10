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

	if users.Age == 12 {
		fmt.Println("users", users.Age, users.Name, users.Salary)
	} else {
		fmt.Println("error")
	}

}
