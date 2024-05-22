// go run .\main.go .\error-handling.go

package main

import (
	"fmt"
)

type specialErr struct {
	msg string
	code int
}

func (specialErr) Error() string {
	return "random error"
}

func ErrorHandling() {
	res, err := returnError(false)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// type error interface {
	// Error() string
// }
// The error body just has to implement a function called Error that returns a string. This is Duck Typing in Go. Doesn't have to look alike, just has to do the same

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", specialErr{ "special lol", 69 }
	}

	return "no error", nil
}