package main

import (
	"fmt"
	"todo_app/pkg/reader"
)

func main() {
	input, err := reader.ReadInput()
	if err != nil {
		panic("O CHUJ")
	}
	fmt.Println(input)
}
