package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/JhonCodeU/go-cli-crud/tasks"
)

func main() {

	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	// [path of program, command]
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("what is the task name? ")
		name, _ := reader.ReadString('\n')
		strings.TrimSpace(name)

		fmt.Println("What is the task description? ")
		description, _ := reader.ReadString('\n')
		strings.TrimSpace(description)

		tasks = task.AddTask(tasks, name, description)
		task.SaveTask(file, tasks)
		fmt.Println("Task added successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Please provide a valid number")
			return
		}

		tasks = task.DeleteTask(tasks, id)
		task.SaveTask(file, tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Please provide a valid number")
			return
		}

		tasks = task.CompleteTask(tasks, id)
		task.SaveTask(file, tasks)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: go-cli-crud [list|add|complete|delete]")
}
