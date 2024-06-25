package main

import (
	"fmt"
)

type Task struct {
	Name        string
	Description string
	Completed   bool
}

var taskList []Task

func addTask(name string, desc string) {

	newTask := Task{

		Name:        name,
		Description: desc,
	}

	taskList = append(taskList, newTask)

	fmt.Println("Task added")
	fmt.Println()

}

func viewTasks() {

	if taskList == nil {

		fmt.Println("Task list empty")
		fmt.Println()

		return
	}

	for _, task := range taskList {

		i := 1

		fmt.Printf("\n------Task Number %d------\n", i)
		fmt.Printf("Name of the task: %s\n", task.Name)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Is the task completed?: ")

		if task.Completed {

			fmt.Printf("Yes\n")
			fmt.Println()

		} else {

			fmt.Printf("No\n")
			fmt.Println()
		}

		i += 1

	}
}

func completeTask(index int) {

	taskList[index].Completed = true

	fmt.Printf("The task named %s has been marked completed!", taskList[index].Name)
	fmt.Println()

}

func main() {

	fmt.Println("Welcome to my Task Management System")

	var choice int

	fmt.Println("\nTask Management System")

	for {

		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")
		fmt.Println()
		fmt.Print("Enter your choice: ")
		fmt.Println()

		fmt.Scan(&choice)

		switch choice {

		case 1:

			var name, description string
			fmt.Print("Enter task name: ")
			fmt.Scan(&name)
			fmt.Print("Enter task description: ")
			fmt.Scan(&description)
			addTask(name, description)

		case 2:
			viewTasks()

		case 3:
			var index int
			fmt.Print("Enter task number to complete: ")
			fmt.Scan(&index)
			completeTask(index - 1)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}

}
