package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func slicesAndMaps() {
	slice := []int{ 1, 2, 3 }
	slice2 := []int{ 4, 5 }
	slice = append(slice, slice2...)
	fmt.Println(slice)

	dict := map[int]string{ 23: "random", 35: "haha" }
	var val, exists = dict[23]

	if exists {
		fmt.Println(val)
	} else {
		fmt.Println("Invalid key")
	}
}

func stringManipulation() {
	thaiString := "สวัสดี"

	// Thai characters takes/starts at 3.5k+ bytes
	for idx, char := range thaiString {
		fmt.Printf("%v %v\n", idx, char) // This prints out their unicode representation
	}

	// thaiStringButIndexable := []rune("สวัสดี")

	strSlice := []string{"ส", "วั", "ส", "ดี"}
	var strBuilder strings.Builder

	for idx := range strSlice {
		strBuilder.WriteString(strSlice[idx])
	}

	fmt.Println(strBuilder.String())
}

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type elecEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e elecEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}

func typesAndStuff() {
	myEngine := gasEngine{ 25, 15 }
	fmt.Println(myEngine.milesLeft())

	canMakeIt(myEngine, 99)
	canMakeIt(elecEngine{ 123, 123 }, 99)

	// anonStruct := struct{
	// 	age int
	// 	height int
	// }{ 12, 111 }
}

func pointers() {
	var p *int32 = new(int32)

	fmt.Printf("%v %v\n", *p, p)
	*p = 10
	fmt.Printf("%v %v", *p, p)

}

var dbData = []string{ "id1", "id2", "id3", "id4", "id5" }
var results = []string{}
var waitGroup = sync.WaitGroup{}

func dbCall(i int) {
	// simulates db call delay
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
	results = append(results, dbData[i])
	waitGroup.Done()
}

func main() {
	fmt.Println()

	// typesAndStuff()
	// pointers()

	t0 := time.Now()

	for i := 0; i < len(dbData); i += 1 {
		waitGroup.Add(1)
		go dbCall(i)
	}

	waitGroup.Wait()
	fmt.Println(results)
	fmt.Printf("Total execution time: %v", time.Since(t0))
}