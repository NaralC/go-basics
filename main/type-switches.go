// go run .\main.go .\type-switches.go

package main

import "fmt"

type star struct {
	owner string
}

type sword struct {
	owner string
}

func (star) attack() {
	fmt.Println("pew pew")
}

func (star) pickUp() {
	fmt.Println("got a star!")
}

func (sword) attack() {
	fmt.Println("chwinnnngg")
}

type weapon interface {
	attack()
}

func TypeSwitches()  {
	weapons := []weapon{
		star{"naral"},
		sword{"not naral"},
	}

	fmt.Println(weapons)

	for _, item := range weapons {
		item.attack()

		switch weapon.(type) {
		case star:
			myStar := weapon.(star)
			myStar.pickUp()
		case sword:
			mySword := weapon.(sword)
		}

	}
}
