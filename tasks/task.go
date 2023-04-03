package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ListTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
	}

	for _, task := range tasks {

		status := ""
		if task.Completed {
			status = "✔"
		} else {
			status = ""
		}

		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Name)
	}

	return tasks
}

func AddTask(tasks []Task, name string, description string) []Task {
	task := Task{
		ID:          GetNextID(tasks),
		Name:        name,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, task)

	return tasks
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	writer, err := file.Write(bytes)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved %d bytes to file\n", writer)
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}

	return tasks
}

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
