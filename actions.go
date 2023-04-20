package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTask(c *cli.Context) error{
	str := c.Args().First()
	fmt.Println(str)
	if str == "" {
		return errors.New("cannot add an empty task")
	}
	
	task := &Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      str,
		Completed: false,
	}
	
	return CreateTask(task)

}