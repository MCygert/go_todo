package repo

import "todo_app/pkg/reader"

type TaskRepository interface {
	Save() error
	Read() reader.Task
}
