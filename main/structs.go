// go run .\main.go .\structs.go

package main

import "fmt"

func Structs()  {
	// Expose attributes by capitalizing them
	type ninja struct {
		name string
		weapons []string
		level int
	}

	wallace := ninja{
		name: "Wallace",
		weapons: []string{"Mace", "Sword"},
		level: 3,
	}

	wallace.weapons = append(wallace.weapons, "Gun")
	fmt.Println(wallace.weapons)

	type dojo struct {
		name string
		ninja ninja
	}

	myDojo := dojo{
		"myDojo",
		wallace,
	}

	fmt.Println(myDojo.ninja.name)

	type addressedDojo struct {
		name string
		ninjaPtr *ninja
	}

	dojo2 := addressedDojo{
		"ptrDojo",
		&wallace,
	}

	dojo2.ninjaPtr.name = "not wallace"
	fmt.Println(dojo2.ninjaPtr.name)

	johnnyPtr := new(ninja)
	johnnyPtr.name = "johnny haha"
	fmt.Println(johnnyPtr)

	naral := intern{
		name: "naral",
		level: 3,
	}
	naral.sayHello()
	naral.sayName()
}

type intern struct {
	name string
	level int
}

func (intern) sayHello() {
	fmt.Println("gimme return offer!11111!!!11!")
}

func (n intern) sayName() {
	fmt.Println("I am " + n.name)
}
