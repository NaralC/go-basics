// go run .\main.go .\text-template.go

package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name string
	Secret string
}

func TextTemplate()  {
	secret := secret{ "Wallace", "I am a ninja" }
	templatePath := `E:\Development\go-basics\main\text-template.go`
	
	t, err := template.New("template.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)

	if err != nil {
		fmt.Println(err)
	}
}
