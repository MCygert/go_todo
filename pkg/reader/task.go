package reader

import "time"

type DomainTask struct {
	_id string
	Task
}
type Task struct {
	Task      string
	AddedDate time.Time
	Expired   bool
}
