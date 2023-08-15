package main

import (
	"flag"
	"fmt"
)

type CommandFlags struct {
	Add    string
	Delete int
	List   bool
}

func NewCommand() *CommandFlags {

	// Below code uses flags to parse command line arguments. But it can be also be done using os.Args
	// args := os.Args

	cFlags := CommandFlags{}

	flag.StringVar(&cFlags.Add, "add", "", "Add a todo")
	flag.IntVar(&cFlags.Delete, "delete", -1, "Delete a todo with id")
	flag.BoolVar(&cFlags.List, "list", false, "List todo")

	flag.Parse()

	return &cFlags

}

func (cFlags *CommandFlags) Execute(todo *Todos) {

	switch {
	case cFlags.List:
		todo.Print()
	case cFlags.Add != "":
		todo.Add(cFlags.Add)
	case cFlags.Delete != -1:
		todo.Delete(cFlags.Delete)

	default:
		fmt.Println("Invalid command")
	}

}
