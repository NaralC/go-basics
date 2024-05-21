// go run .\main.go .\constants.go

package main

import (
	"fmt"
	"time"
)

func Constants()  {
	const (
		a int = 1
		b  = 2
		c int = 3
		d  = 4
	)

	const (
		zero int = iota + 4
		one
		two
	)

	fmt.Println(zero, one, two)
	fmt.Printf("%T", time.December)
}