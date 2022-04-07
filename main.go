package main

import "todo_app/cli"

func main() {
	input, err := cli.ReadInput()
	if err != nil {
		panic("O CHUJ")
	}
}
