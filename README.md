# Task Management System

## TODO

- [X] step1: connecting to MongoDB
- [X] step2: creating a CLI program
- [X] step3: creating a task
- [X] step4: show all tasks
- [X] step5: complete a task
- [X] step6: displaying pending tasks only
- [X] step7: displaying finished tasks only
- [X] step8: deleting a task

## Third party packages

1. `go get go.mongodb.org/mongo-driver`
2. `go get github.com/urfave/cli/v2`
3. `go get gopkg.in/gookit/color.v1`
4. `go get github.com/joho/godotenv`

## Complete a task

For completing a task we need to grab a string value from command-line in order to identify the task we wanted to update, then with some efforts over creating an appropriate query we do the trick.

## List all completed tasks

Just to be clear you only need to know how to use a good query statement and all the rest is the same code as previous parts.

``` go
filter := bson.D {
    primitive.E{Key: "completed", Value: true},
}
```

## Deleting a task

```go
filter := bson.D{primitive.E{Key: "text", Value: text}}
```

## Output

![trailer](./docs/output.gif)
