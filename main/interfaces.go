// go run .\main.go .\interfaces.go

package main

import "fmt"

type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon) {
	fmt.Println("bruh")
	weapon.attack()
}

type ninjaStar struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

type ninjaSword struct {
	owner string
}

func (ninjaSword) attack() {
	fmt.Println("Swinging sword")
}

func Interfaces()  {
	weapons := []ninjaWeapon{ ninjaStar{ "naral" }, ninjaSword{ "naral" } }

	for _, weapon := range weapons {
		attack(weapon)
		// Why not weapon.attack()?
	}
}