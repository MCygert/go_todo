package reader

import "time"

type Task struct {
	Task      string
	AddedDate time.Time
	Expired   bool
}
