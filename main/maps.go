// go run .\main.go .\maps.go

package main

import "fmt"

func Maps()  {
	// var myMap map[string]string

	// m := make(map[string]string, 5)

	x := map[string]string{ "Wallace": "ninja" }
	x["Johnny"] = "student"
	fmt.Println(x)
	delete(x, "Johnny")
	x["Wallace"] += " not capping"
	fmt.Println(x)

	for key, val := range x {
		fmt.Println(key, val)
	}

	_, exists := x["qwf"]
	fmt.Println(exists)
}