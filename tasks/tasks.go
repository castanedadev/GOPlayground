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

func main() {
	fmt.Println("=== Tasks Summary ===")
}
