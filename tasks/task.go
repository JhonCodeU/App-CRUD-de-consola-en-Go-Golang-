package task

import "fmt"

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
			status = "✘"
		}

		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Name)
	}

	return tasks
}
