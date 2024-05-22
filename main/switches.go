// go run .\main.go .\switches.go

package main

import "fmt"

func Switches()  {
	venues := []string{"Home", "Work", "School", "Bar", "Gym"}

	for _, val := range venues {
		switch val {
		case "Home":
			greeting("Mom me hungry")
		case "Work", "School":
			greeting("Weather's great XD")
		default:
			greeting("bruh")
		}
	}
}

func greeting(greeting string)  {
	fmt.Println(greeting)
}