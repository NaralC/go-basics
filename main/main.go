package main

import (
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
	Constants()
}