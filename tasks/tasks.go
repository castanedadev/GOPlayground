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
	fmt.Println("Completed Tasks...")
	completedTasksCount := 0
	for _, task := range t.tasks {
		if task.isCompleted {
			fmt.Println("Name => ", task.name)
			fmt.Println("Description => ", task.description)
			completedTasksCount++
		}
	}

	if completedTasksCount <= 0 {
		fmt.Println("No completed tasks yet!")
	}
}

func (t *taskList) printTasksInProgress() {
	fmt.Println("Tasks In Progress...")
	for _, task := range t.tasks {
		if !task.isCompleted {
			fmt.Println("Name =>", task.name)
		}
	}
}

func main() {
	tfr1 := &task{
		name:        "Tâche 1",
		description: "Ma premiere tâche",
	}
	tfr2 := &task{
		name:        "Tâche 2",
		description: "Ma deuxieme tâche",
	}
	tfr3 := &task{
		name:        "Tâche 3",
		description: "Ma troisième tâche",
	}

	frenchTasks := &taskList{
		tasks: []*task{
			tfr1, tfr2,
		},
	}

	frenchTasks.addToList(tfr3)
	frenchTasks.tasks[0].markAsComplete()

	// Introduce map to handle multiple task lists (E.G. by subject, by user, etcetera)
	tasksMap := make(map[string]*taskList)
	tasksMap["French"] = frenchTasks

	tEnglish1 := &task{
		name:        "Task 1",
		description: "My first task",
	}
	tEnglish2 := &task{
		name:        "Task 2",
		description: "My second task",
	}
	tEnglish3 := &task{
		name:        "Task 3",
		description: "My third task",
	}

	englishTasks := &taskList{
		tasks: []*task{
			tEnglish1, tEnglish2, tEnglish3,
		},
	}

	tasksMap["English"] = englishTasks

	// Print task elements by retrieving list from map

	fmt.Println("=== French Tasks Summary ===")
	fmt.Print("\n")
	tasksMap["French"].printTasksInProgress()
	fmt.Print("\n")
	tasksMap["French"].printCompletedTasks()

	fmt.Print("\n")
	fmt.Println("=== English Tasks Summary ===")
	fmt.Print("\n")
	tasksMap["English"].printTasksInProgress()
	fmt.Print("\n")
	tasksMap["English"].printCompletedTasks()
}
