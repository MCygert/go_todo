package main

import (
	"fmt"
	"todo_app/internal/repo"
	"todo_app/pkg/reader"
)

func main() {
	input, err := reader.ReadInput()
	if err != nil {
		panic(err)
	}
	err = repo.Init()
	if err != nil {
		panic(err)
	}
	repo.SaveTask(input)
	fmt.Println(input)
}
