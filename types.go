package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gookit/color.v1"
)

type Task struct {
	// we use primitive since MongoDB used `objectID`
	// for `_id` field by default.
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

// create a function that receives an instance of Task
// and saves it in the database.
func CreateTask(task *Task) error {

	_, err := collection.InsertOne(ctx, task)
	return err
}

func GetAll() ([]*Task, error) {
	var tasks []*Task

	filter := bson.D{{}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &task)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}
	
	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}


	return tasks, nil
}

func printTasks(tasks []*Task) {
	for i, v := range tasks {
		if v.Completed {
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else {
			color.Yellow.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}
