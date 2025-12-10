package customerrorhandle

// Structs → replace classes

// Composition → replaces inheritance

// Interfaces → provide polymorphism

// Methods → attach behavior to structs
import (
	"errors"
	"fmt"
)

type Message struct {
	Messages string

	Code int
}

func (e *Message) Error() string {

	return fmt.Sprintf("Custom Error", e.Code, e.Messages)

}

var ErrorNotFound = errors.New("Item Not Found")

func FindItem(num int) (string, error) {

	if num == 0 {

		return "", &Message{Messages: "Invalid ", Code: 100}
	}

	if num == 1 {

		return "", ErrorNotFound

	}

	return "Found item", nil
}
