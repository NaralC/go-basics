package main

import (
	"errors"
	"fmt"
	"math"
)

func variablesAndTypes() {
	// var name type = expression
	// defaults to 0, "", false, null if unassigned
	var integer int
	var integer1, integer2 = 2, "string"

	// name := expression
	goIsAwesome, javaIsAwesome := true, false

	fmt.Println(integer)
	fmt.Println(integer1, integer2)
	fmt.Println(goIsAwesome, javaIsAwesome)
}

func pointers()  {
	// var x int = 1
	// var p *int = &x
	x := 1
	p := &x

	fmt.Println(p)
	fmt.Println(*p)
}

func types() {
	type farenheit int
	type celsius int

	var f farenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}

func numbers()  {
	var x int16 = 1
	y := 15.289
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(y + 2)

	fmt.Println(math.Ceil(3.99))
	fmt.Println(math.Min(3.99, 3892))
	fmt.Println(math.Abs(3 - 7))
}

func main()  {
	// variablesAndTypes()
	// pointers()
	// types()

	// {
	// 	var diff = 2	
	// 	fmt.Println(diff)
	// }

	// var diff = 1
	// fmt.Println(diff)
	// fmt.Println(global)
	// numbers()
	// ExportedFunction()
	// Strings()
	// StringBuilder()
	// Constants()
	// Loops()
	// Arrays()
	// Structs()
	// Maps()
	// Slices()
	// Switches()
	// TextTemplate()
	// JSON()
	// variadic(1,2,3,4,5,6)
	// sum(1, 5)
	// fmt.Println(nameLength("naral"))

	// res, err := sayHello("naral")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	// // fibonacci(10)
	// ptrNumber := 5
	// usePtr(&ptrNumber)
	// fmt.Println(ptrNumber)
	// AdvFuncs()
	Interfaces()
}

func testFunc() {
	fmt.Println("Yes")
}

// As many ints as you want
func variadic(numbers ...int) {
	for _, integer := range numbers {
		fmt.Println(integer)
	}
}

func sum(a int, b int) int {
	return a + b
}

func nameLength(name string) (string, int) {
	return name, len(name)
}

func sayHello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Empty name btw")
	} else {
		return "Hello " + name, nil
	}
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i - 1) + fibonacci(i - 2)
	}
}

func usePtr(i *int) {
	*i = *i + 1
}
