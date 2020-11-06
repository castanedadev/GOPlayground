package main

import "fmt"

// Structs definitions
type task struct {
	name        string
	description string
	isCompleted bool
}
type taskList struct {
	tasks []*task
}

// Consumer functions
func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteFromList(index int) {
	t.tasks = append(t.tasks[:index])
}

func (t *task) updateDescription(desc string) {
	t.description = desc
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) markAsComplete() {
	t.isCompleted = true
}

func (t *taskList) printCompletedTasks() {
	for _, task := range t.tasks {
		if task.isCompleted {
			fmt.Println("Name => ", task.name)
			fmt.Println("Description => ", task.description)
		}
	}
}

func (t *taskList) printTasksInProgress() {
	for _, task := range t.tasks {
		if !task.isCompleted {
			fmt.Println("Name =>", task.name)
		}
	}
}

func main() {
	fmt.Println("=== Tasks Summary ===")
	task1 := &task{
		name:        "Task 1",
		description: "Ma premiere tâche",
	}
	task2 := &task{
		name:        "Task 2",
		description: "Ma deuxieme tâche",
	}
	task3 := &task{
		name:        "Task 3",
		description: "Ma troisième tâche",
	}

	taskList := &taskList{
		tasks: []*task{
			task1, task2,
		},
	}

	taskList.addToList(task3)
	taskList.printCompletedTasks()
	taskList.tasks[0].markAsComplete()

	fmt.Print("\n")
	fmt.Println("Tasks In Progress...")
	taskList.printTasksInProgress()

	fmt.Print("\n")
	fmt.Println("Completed Tasks...")
	taskList.printCompletedTasks()
}
