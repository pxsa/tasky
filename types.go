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
	filter := bson.D{}
	return getAllByFilter(filter)
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

// complete a task
func CompleteTask(taskName string) error{
	filter := bson.D{primitive.E{Key: "text", Value: taskName}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &Task{}
	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}

// get all penidng tasks
func GetAllPending() ([]*Task, error) {
	filter := bson.D {
		primitive.E{Key:"completed", Value: false},
	}
	return getAllByFilter(filter)
}

// Get all Tasks by filter
func getAllByFilter(filter bson.D) ([]*Task, error) {
	var tasks []*Task

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cursor.Next(ctx) {
		var task Task
		err = cursor.Decode(&task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &task)
	}
	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		return tasks, err
	}
	
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

// Get all completed tasks
func GetAllCompleted() ([]*Task, error) {
	filter := bson.D {
		primitive.E{Key: "completed", Value: true},
	}

	return getAllByFilter(filter)
}
