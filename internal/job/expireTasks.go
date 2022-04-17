package job

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
	"todo_app/internal/repo"
)

func updateExpiredTasks() {
	notExpired, err := repo.FindAllNotExpired()
	now := time.Now()
	if err != nil {
		log.Fatal(err)
	}
	numberOfExpired := 0
	for _, expired := range notExpired {
		timeBefore := expired.AddedDate.Add(time.Hour * 24)
		if timeBefore.After(now) {
			numberOfExpired++
			expired.Expired = true
			repo.UpdateTask(expired)
		}
	}
	log.Println(numberOfExpired)
}

func StartExpiringTasksJob() {
	c := cron.New()
	_, err := c.AddFunc("@every 1m", updateExpiredTasks)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
