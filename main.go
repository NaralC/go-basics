package main

import "fmt"

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

func main()  {
	// variablesAndTypes()
	// pointers()
	types()
}