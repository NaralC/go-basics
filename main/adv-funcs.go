// go run .\main.go .\adv-funcs.go

package main

import "fmt"

func AdvFuncs()  {
	fmt.Printf("%T\n", funcAsParam(5, func(i int) int {
		return i * 2 // Anonymous function that doubles the input
	}))

	defer first()
	second()
	second()
	second()

	// anonymous function
	func() {
		fmt.Println("anon func lol")
	}()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func funcAsParam(i int, f func(int) int) int {
	return f(i)
}
