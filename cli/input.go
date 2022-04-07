package cli

import (
	"flag"
	"fmt"
	"time"
)

type Task struct {
	Task      string
	AddedDate time.Time
}

func ReadInput() (Task, error) {
	var task string
	date := time.Now()

	flag.StringVar(
		&task,
		"t",
		"",
		"Specify task")

	flag.Parse()
	if task == "" {
		return Task{}, fmt.Errorf("pass task")
	}
	return Task{task, date}, nil

}
