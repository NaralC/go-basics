// go run .\main.go .\loops.go

package main

import "fmt"

func Loops()  {

	for i := 0; i < 10; i += 1 {
		fmt.Println(i)
	}

	str := "I love TypeScript"

	// for idx, char := range str {
	// 	fmt.Printf("%d %c", idx, char)
	// }

	for _, char := range str {
		if char == 'T' {
			break
		}

		fmt.Printf("%c\n", char)
		continue
	}

	// Labels
	iForLoop:
	for i := 0; i < 5; i += 1 {
		for j := 0; j < 5; j += 1 {
			if j == 3 {
				break iForLoop
			}

			fmt.Printf("%d %d\t", i, j)
		}

		fmt.Println()
	}
}