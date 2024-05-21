// go run .\main.go .\strings.go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Strings() {
	s := "hello world"

	fmt.Println(len(s))
	fmt.Println(s[0]) // Prints out ASCII value
	fmt.Printf("%c\n", s[0]) // Prints out actual string
	fmt.Println(s[0:6]) // String slicing
	fmt.Println(s + " again") // String concatenation
	fmt.Println(s + "\b arstarst") // bruh
	
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Println("ฮัลโหล เวิล์ด") // Non-ASCII support
}

func StringBuilder() {
	// a := "henlo world"

	// fmt.Println(strings.ToUpper((a)))
	// fmt.Println(strings.HasPrefix(a, "henlo"))
	// fmt.Println(strings.Replace(a, "world", "replaced", 1))
	// fmt.Println(strings.Replace(a + " world", "world", "replaced", 2))
	 
	var sb strings.Builder
	sb.WriteString("Hennnlo")
	fmt.Println(sb.String())
	fmt.Println(sb.Cap(), sb.Len())
	sb.Grow(2)
	fmt.Println(sb.Cap(), sb.Len())
	sb.Reset()
	sb.Write([]byte{ 66, 66, 65 })
	fmt.Println(strings.ToLower(sb.String()))

	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y)
	z, _ := strconv.Atoi((y)) 
	fmt.Println(z)
}