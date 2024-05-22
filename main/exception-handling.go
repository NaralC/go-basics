// go run .\main.go .\error-handling.go

package main

import "fmt"

func ExceptionHandling()  {
	defer safeExit()

	const one = 2

	if one != 1 {
		panic("One is not 1????")
	}
}

func safeExit() {
	r := recover()

	if r != nil {
		fmt.Println("Panic's recovered!")
	}
}
