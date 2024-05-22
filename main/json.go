// go run .\main.go .\json.go

package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name string `json:"full_name"`
	Level int
}

func JSON()  {
	jsonStr := `{ "full_name": "Somjai", "level": 1 }`

	var wallace Ninja
	err := json.Unmarshal([]byte(jsonStr), &wallace)
	if err != nil {
		fmt.Println("woops")
	}
	fmt.Println(wallace)

	sTo, err := json.Marshal(wallace)
	fmt.Println(string(sTo))
}
