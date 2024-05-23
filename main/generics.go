// go run .\main.go .\generics.go

package main

import "fmt"

type superFloat float64

type minTypes interface {
	~float64 | int
}

func Generics() {
	fmt.Println(minInt(1, 5))
	fmt.Println(minFloat64(1.45, 5.11))
	// fmt.Println(min[int](4, 3))

	var sf superFloat = 0.3
	fmt.Println(min(sf, 3))
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func minFloat64(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func min[T minTypes](a T, b T) T {
	return a
}
