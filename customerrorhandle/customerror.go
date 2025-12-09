package customerrorhandle

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
