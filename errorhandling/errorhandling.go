package errorhandling

import (
	"errors"
	"fmt"
)

func GetError() {

	fmt.Println("hlo world")

}

// return type string and int
func DoSomething() (string, error) {

	var s string

	s = "shakti"
	if true {
		return s, errors.New("something went wrong")
	}
	return "success", nil
}
