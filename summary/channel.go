package main

import "fmt"

func Channels() {
	var channel = make(chan int)
	go process(channel)
	fmt.Println(<- channel)
}

func process(c chan int) {
	c <- 123
}
