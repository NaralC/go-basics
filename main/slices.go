// go run .\main.go .\slices.go

package main

import "fmt"

func Slices()  {
	fixed := [...]int{0, 1, 2}
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)

	fmt.Println(fixed, slice)

	b := make([]int, 4) // all values default to 0
	fmt.Println(b)
}