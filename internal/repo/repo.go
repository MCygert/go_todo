package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"todo_app/pkg/reader"
)

var db *mongo.Client
var taskCollection *mongo.Collection

type TaskRepository interface {
	Save() error
	Read() reader.Task
}

func Init() error {
	var err error
	if db == nil {
		db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/tasks"))
	}
	if err != nil {
		return err
	}
	taskCollection = db.Database("tasks").Collection("tasks")
	return nil
}

func AllTasks() ([]reader.Task, error) {
	cursor, err := taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	var results []reader.Task
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func SaveTask(task reader.Task) {
	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		panic(err)
	}
}

func GetAllNotExpired() ([]reader.Task, error) {
	cursor, err := taskCollection.Find(context.TODO(), bson.M{"expired": false})
	if err != nil {
		log.Fatal(err)
	}
	var results []reader.Task
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
