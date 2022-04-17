package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo_app/internal/job"
	"todo_app/internal/repo"
)

func main() {
	err := repo.Init()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/tasks", tasksIndex)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	job.StartExpiringTasksJob()
}

func tasksIndex(w http.ResponseWriter, _ *http.Request) {
	tasks, err := repo.FindAllTasks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, tsk := range tasks {
		fmt.Println(tsk)
	}

}
