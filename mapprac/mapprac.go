package mapprac

import (
	"fmt"
)

var (
	pracmap = make(map[string]int, 2)
)

type User struct {
	Name          string
	Age           int
	ContactNumber int64
}

func MapPrac() {

	fmt.Println("map")

	var a []int

	a = []int{1, 2, 3, 4}

	a = append(a, 1, 2, 3)

	// 	Maps are used to store data values in key:value pairs.

	// Each element in a map is a key:value pair.

	for _idex := range a {

		fmt.Println(a[_idex])
	}

	prac := map[string]string{"name": "shakti"}

	fmt.Println(prac)

	pracmap["age"] = 12

	fmt.Println(pracmap)

	var users []User

	mapuser := make([]User, 0, 3)

	perac := User{Name: "shakti", Age: 12, ContactNumber: 9841112420}

	users = append(users, perac)

	mapuser = append(mapuser, users...)

	for _, index := range mapuser {

		fmt.Println(index.Age, "\t", index.ContactNumber)
	}

}
