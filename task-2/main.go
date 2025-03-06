package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Id        int
	Name      string
	Completed bool
}

var taskList = []Task{}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark a task as completed")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			markTaskAsCompleted(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	taskList = append(taskList, Task{
		Id:        len(taskList) + 1,
		Name:      taskName,
		Completed: false,
	})
	fmt.Println("Task added successfully")
}

func listTasks() {
	fmt.Println("List tasks")

	if len(taskList) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range taskList {
		status := "" // Default to not done
		if task.Completed {
			status = "✓" // Mark as completed
		}
		fmt.Printf("%d. %s %s\n", task.Id, task.Name, status)
	}
}

func markTaskAsCompleted(reader *bufio.Reader) {
	fmt.Print("Enter task ID to mark as completed: ")
	taskId, _ := reader.ReadString('\n')
	taskId = strings.TrimSpace(taskId)

	for i, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == taskId {
			taskList[i].Completed = true
			fmt.Println("Task marked as completed")
			return
		}
	}

	fmt.Println("Task not found")
}

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to delete: ")
	taskId, _ := reader.ReadString('\n')
	taskId = strings.TrimSpace(taskId)

	for i, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == taskId {
			taskList = append(taskList[:i], taskList[i+1:]...)
			fmt.Println("Task deleted successfully")
			return
		}
	}

	fmt.Println("Task not found")
}
