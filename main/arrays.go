// go run .\main.go .\arrays.go

package main

import "fmt"

func Arrays()  {
	var a [3]int
	fmt.Println(a)
	a[0] = 150
	fmt.Println(a)
	
	aCopy := a // This is a copy, not a ref
	fmt.Println(aCopy == a)
	aCopy[len(aCopy) - 1] = 22
	fmt.Println(aCopy)

	b := [4]int{ 1, 2, 3, 4 }
	fmt.Println(b)

	c := [...]int{ 0: 4, 4: 1 } // 0th = 4, 4th = 1
	fmt.Println(c)

	for _, ele := range b {
		fmt.Println(ele)
	}

	strArr := [...]string{ "str1", "str2" }
	fmt.Println(strArr)

	twoDimensional := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDimensional)
}