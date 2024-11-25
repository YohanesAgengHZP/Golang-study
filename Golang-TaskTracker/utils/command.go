package utils

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"tasktracker/controllers"
)

type CommandFlags struct {
	// ... fileds for command ...
	Add string
	Delete int
	Edit string
	Toogle int
	List bool
}

func NewCommandFlags() *CommandFlags {
	commndFlag := CommandFlags{}
	flag.StringVar(&commndFlag.Add, "add", "", "Add task to list of tasks")
	flag.IntVar(&commndFlag.Delete, "delete", -1, "Specify task to delete")
	flag.StringVar(&commndFlag.Edit, "edit", "", "Edit task in the list of tasks")
	flag.IntVar(&commndFlag.Toogle, "toogle", -1, "Status of the task")
	flag.BoolVar(&commndFlag.List, "list", false, "List tasks")

	flag.Parse()
	return &commndFlag
}

func (commandFlag *CommandFlags) Execute(todos *controllers.Todos) error {
	switch {
	case commandFlag.List:
		todos.Print()

	case commandFlag.Add != "":
		todos.Add(commandFlag.Add)

	case commandFlag.Edit != "":
		parts := strings.SplitN(commandFlag.Edit, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid edit command format. Use 'id:new_title'")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid index in edit command: %v", err)
		}
		err = todos.Edit(index, parts[1], false)
		if err != nil {
			return err
		}

	case commandFlag.Toogle!= -1:
		err := todos.Toggle(commandFlag.Toogle)
		if err != nil {
			return err
		}

	case commandFlag.Delete != -1:
		err := todos.Delete(commandFlag.Delete)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid command")
	}

	return nil
}
