package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// define some package level variables
var (
	collection *mongo.Collection

	// like a timeout or deadline that indicates when an operation
	// should stop running and return. It helps to prevent performance
	// degradation on production sytems.
	ctx = context.TODO()
)

// a task manager system written in go and using mongoDb
func main() {
	app := &cli.App{
		Name:     "tasker",
		Usage:    "A simple CLI program to manage your tasks",
		Commands: []*cli.Command{

			// add a new task
			{
				Name:    "add",
				Aliases: []string{"-a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					taskTitle := c.Args().First()
					if taskTitle == "" {
						return errors.New("cannot add an empty task")
					}
					newTask := &Task{
						ID: primitive.NewObjectID(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Text: taskTitle,
						Completed: false,
					}
					return CreateTask(newTask)
				},
			},

			// list all tasks
			{
				Name: "all",
				Aliases: []string{"-l"},
				Usage: "list all tasks",
				Action: func(c *cli.Context) error {
					tasks, err := GetAll()
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
							return nil
						}
						return err
					}
					printTasks(tasks)
					return nil
				},
			},

			// complete a task
			{
				Name: "done",
				Aliases: []string{"-d"},
				Usage: "complete a task",
				Action: func(c *cli.Context) error {
					taskName := c.Args().First()
					return CompleteTask(taskName)
				},
			},

			// list pending tasks only
			{
				Name: "pending",
				Aliases: []string{"p"},
				Usage: "list pending tasks",
				Action: func(ctx *cli.Context) error {
					tasks, err :=  GetAllPending()
					printTasks(tasks)
					return err
				},
			},

			// list all completed tasks
			{
				Name: "completed",
				Aliases: []string{"c"},
				Usage: "list all completed tasks",
				Action: func(ctx *cli.Context) error {
					tasks, err :=  GetAllCompleted()
					printTasks(tasks)
					return err
				},
			},

			// deleting a task
			{
				Name: "delete",
				Aliases: []string{"D"},
				Usage: "delete a task",
				Action: func(ctx *cli.Context) error {
					taskName := ctx.Args().First()
					return DeleteTask(taskName)
				},
			},

		},
	}


	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	// options.ClientOption used to set the connection string
	// and other driver settings.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	// uri := "mongodb://localhost:27017/"

	
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// let's ensure that out MongoDB server was found
	// and connected successfully using Ping method.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// The Database and Collection types can be
	// used to access the database
	collection = client.Database("task-manager").Collection("tasks")
}
