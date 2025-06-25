package main

import (
	"fmt"
	"os"
)

// Task represents a single todo item
type Task struct {
	ID          int
	Description string
	Completed   bool
}

// TodoList manages a collection of tasks
type TodoList struct {
	Tasks  []Task
	NextID int
}

// NewTodoList creates and initializes a new todo list
func NewTodoList() *TodoList {
	return &TodoList{
		Tasks:  make([]Task, 0),
		NextID: 1,
	}
}

// AddTask adds a new task to the list
func (l *TodoList) AddTask(description string) Task {
	task := Task{
		ID:          l.NextID,
		Description: description,
		Completed:   false,
	}
	
	l.Tasks = append(l.Tasks, task)
	l.NextID++
	
	return task
}

// ListTasks prints all tasks in the list
func (l *TodoList) ListTasks() {
	if len(l.Tasks) == 0 {
		fmt.Println("No tasks in the list!")
		return
	}
	
	fmt.Println("Tasks:")
	fmt.Println("------")
	
	for _, task := range l.Tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Description)
	}
}

// CompleteTask marks a task as completed
func (l *TodoList) CompleteTask(id int) error {
	for i, task := range l.Tasks {
		if task.ID == id {
			l.Tasks[i].Completed = true
			return nil
		}
	}
	
	return fmt.Errorf("task with ID %d not found", id)
}

// TODO: Implement DeleteTask function to remove a task by ID

// TODO: Implement a way to save tasks to disk and load them back

func main() {
	// Initialize a new todo list
	todoList := NewTodoList()
	
	// This is a very basic CLI - you'll need to expand it
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	
	command := os.Args[1]
	
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task description")
			printUsage()
			return
		}
		taskDesc := os.Args[2]
		task := todoList.AddTask(taskDesc)
		fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
		
	case "list":
		todoList.ListTasks()
		
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			printUsage()
			return
		}
		
		// TODO: Convert the argument to an integer and handle the error
		// TODO: Call todoList.CompleteTask() with the ID
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add TASK      Add a new task")
	fmt.Println("  todo list          List all tasks")
	fmt.Println("  todo complete ID   Mark task as completed")
	// TODO: Add command for deleting tasks
}
