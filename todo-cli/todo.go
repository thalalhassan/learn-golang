package main

import (
	"fmt"
	"time"

	"github.com/rodaine/table"
)

type Todo struct {
	ID          int
	Task        string
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todo) IsCompleted() bool {
	return !t.CompletedAt.IsZero()
}

func (todos *Todos) Add(task string) {
	newTodo := Todo{
		ID:          time.Now().Nanosecond(),
		Task:        task,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, newTodo)
}

func (todos *Todos) Delete(id int) {
	curTodos := *todos
	index := todos.GetIndex(id)
	if index == -1 {
		fmt.Println("No todo found")
		return
	}

	*todos = append(curTodos[:index], curTodos[index+1:]...)
}

func (todos *Todos) GetIndex(id int) int {
	index := -1
	// check if id exist
	for i, v := range *todos {
		if v.ID == id {
			index = i
			break
		}
	}

	return index
}

func (todos *Todos) MarkDone(id int) {
	curTodos := *todos
	index := todos.GetIndex(id)
	if index == -1 {
		fmt.Println("No todo found")
		return
	}

	if curTodos[index].CompletedAt != nil {
		fmt.Println("Already done")
		return
	}

	completedAt := time.Now()
	curTodos[index].CompletedAt = &completedAt
}

func (todos *Todos) Print() {
	tbl := table.New("ID", "Task", "CreatedAt", "CompletedAt", "Done")

	for _, t := range *todos {
		completedAt := ""
		done := "x"

		if t.CompletedAt != nil {
			done = "/"
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}

		tbl.AddRow(t.ID, t.Task, t.CreatedAt.Format(time.RFC1123), completedAt, done)
	}

	fmt.Println("========================================================")
	tbl.Print()
	fmt.Println("========================================================")

}
