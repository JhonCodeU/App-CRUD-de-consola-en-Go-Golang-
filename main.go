package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
	}
}

func printUsage() {
	fmt.Println("Usage: go-cli-crud [list|add|complete|delete]")
}
