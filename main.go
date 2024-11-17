package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var tasks []Task

func readFile() []Task {
	filepath := "./data.json"
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file '%s': %v", filepath, err)
	}
	defer file.Close() // Ensure file is closed
	fileContent, _ := io.ReadAll(file)
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		log.Fatalf("Failed to decode json %v", err)
	}
	return tasks
}

func saveFile(tasks []Task) {

	filepath := "data.json"

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file '%s': %v", filepath, err)
	}
	defer file.Close() // Ewnsure file is closed

	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		_ = fmt.Errorf("failed to marshal task %w", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		_ = fmt.Errorf("failed to write to file: %w", err)
	}
	fmt.Println("Data successfully written to file")
}

func Add(description string) {
	// get the length of the tasks
	newIndex := len(tasks) + 1
	var newTask Task

	newTask.Id = newIndex
	newTask.Description = description
	newTask.Status = "pending"
	newTask.CreatedAt = time.Now().Format(time.ANSIC)
	newTask.UpdatedAt = time.Now().Format(time.ANSIC)

	tasks = append(tasks, newTask)
	saveFile(tasks)
}

func Update(index int, data ...string) {
	if len(tasks) > index {
		task := tasks[index]

		if len(data) > 0 {
			description := data[0]
			task.Description = description
			task.UpdatedAt = time.Now().Format(time.ANSIC)
			tasks[index] = task
			saveFile(tasks)
			fmt.Println("Task updated successfully", tasks[index])
		}
	} else {
		fmt.Println("Invalid Index")
	}
}

func MarkDone(index int) {
	if len(tasks) > index {
		task := tasks[index]
		task.Status = "done"
		task.UpdatedAt = time.Now().Format(time.ANSIC)
		tasks[index] = task
		saveFile(tasks)
		fmt.Println("Task updated successfully", tasks[index])
	}
}

func MarkInProgress(index int) {
	if len(tasks) > index {
		task := tasks[index]

		task.Status = "in-progress"
		task.UpdatedAt = time.Now().Format(time.ANSIC)

		tasks[index] = task
		saveFile(tasks)
		fmt.Println("Task status updated successfully", tasks[index])
	}
}

func displayTasks(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("%v  %s   %s   %s   %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func FilterTask(filter string) []Task {
	var tempTasks []Task

	for _, task := range tasks {
		if task.Status == filter {
			tempTasks = append(tempTasks, task)
		}
	}
	
	return tempTasks
}

func List(filter ...string) {

	if len(filter) > 0 {

		switch filter[0] {
		case "todo":
			tempTasks := FilterTask("todo")
			displayTasks(tempTasks)
		case "done":
			tempTasks := FilterTask("done")
			displayTasks(tempTasks)
		case "in-progress":
			tempTasks := FilterTask("in-progress")
			displayTasks(tempTasks)
		default:
			displayTasks(tasks)
		}
	}
}

func Delete(index int) {
	if len(tasks) > 0 {
		tasks = append(tasks[:index], tasks[index+1:]...)
	}
	saveFile(tasks)
	fmt.Println("task has been deleted")
}

func main() {
	var (
		index       int
		description string
		filter      string
		status      string
	)

	tasks = readFile()

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listCommand.StringVar(&filter, "filter", "all", "")

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addCommand.StringVar(&description, "description", "", "")

	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	updateCommand.IntVar(&index, "index", 0, "")
	updateCommand.StringVar(&description, "description", "", "")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCommand.IntVar(&index, "index", 0, "")

	markDone := flag.NewFlagSet("mark-done", flag.ExitOnError)
	markDone.IntVar(&index, "index", 0, "")

	markInProgress := flag.NewFlagSet("mark-in-progress", flag.ExitOnError)
	markInProgress.IntVar(&index, "index", 0, "")

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
		Add(description)
	case "delete":
		deleteCommand.Parse(os.Args[2:])
		Delete(index)
	case "update":
		updateCommand.Parse(os.Args[2:])
		Update(index, description, status)
	case "list":
		listCommand.Parse(os.Args[2:])
		List(filter)
	case "mark-done":
		markDone.Parse(os.Args[2:])
		MarkDone(index)
	case "mark-in-progress":
		markInProgress.Parse(os.Args[2:])
		MarkInProgress(index)
	default:
		fmt.Print("Invalid command")
	}

}
