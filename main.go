package main

import (
	"bufio"
	"fmt"
	"os"
)

func ExistCheck() {
	_, err := os.Stat("tasks.json")
	if err != nil {
		file, _ := os.Create("tasks.json")
		file.WriteString(`{"Tasks": []}`)
		file.Close()
	}
}

func main() {
	ExistCheck()
	TaskArr := NewTaskList()
	err := TaskArr.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Use the following:\nadd \"task\"\nupdate <task id> \"new task description\"\ndelete <task id>\nmark-in-progress <task id>\nmark-done <task id>\nlist <todo/in-progress/done>\n\nType your command:")
	r := bufio.NewReader(os.Stdin)

	for {
		command, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Reading error: ", err)
			continue
		}
		args, id, err := ProcessCommand(command)

		if err != nil {
			fmt.Println(err)
			continue
		}
		length := len(args)
		switch {
		case length == 2 && args[0] == "add":
			id, err := TaskArr.NewTask(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = TaskArr.StoreData()
			if err != nil {
				fmt.Println("Saving data error: ", err)
				continue
			}
			fmt.Printf("Task added successfully (ID: %d)\n\n", id)

		case length == 3 && args[0] == "update":
			id, err := TaskArr.Update(id, args[1])
			if err != nil {
				fmt.Println("Updating error: ", err)
				continue
			}
			err = TaskArr.StoreData()
			if err != nil {
				fmt.Println("Saving data error: ", err)
				continue
			}
			fmt.Printf("Updating is succesful (ID: %d)\n\n", id)

		case length == 2 && args[0] == "delete":
			id, err := TaskArr.Delete(id)
			if err != nil {
				fmt.Println("Deleting error: ", err)
				continue
			}
			err = TaskArr.StoreData()
			if err != nil {
				fmt.Println("Saving data error: ", err)
				continue
			}
			fmt.Printf("Deleting is succesful (ID: %d)\n\n", id)

		case length == 2 && args[0] == "mark-in-progress":
			id, err := TaskArr.SetInProgress(id)
			if err != nil {
				fmt.Println("Parsing ID error: ", err)
				continue
			}
			err = TaskArr.StoreData()
			if err != nil {
				fmt.Println("Saving data error: ", err)
				continue
			}
			fmt.Printf("Task with ID: %d is in progress now\n\n", id)

		case length == 2 && args[0] == "mark-done":
			id, err := TaskArr.SetDone(id)
			if err != nil {
				fmt.Println("Parsing ID error: ", err)
				continue
			}
			err = TaskArr.StoreData()
			if err != nil {
				fmt.Println("Saving data error: ", err)
				continue
			}
			fmt.Printf("Task with ID: %d is done now\n\n", id)

		case length <= 2 && args[0] == "list":
			switch len(args) {
			case 2:
				switch args[1] {
				case "in-progress":
					TaskArr.list(status("In progress"))
				case "done":
					TaskArr.list(status("Done"))
				case "todo":
					TaskArr.list(status("To do"))
				}
			default:
				TaskArr.list(status(""))
			}

		default:
			fmt.Print("\nUnknown command\nUse the following:\nadd \"task\"\nupdate <task id> \"new task description\"\ndelete <task id>\nmark-in-progress <task id>\nmark-done <task id>\nlist <todo/in-progress/done>\n\n")
		}
	}
}
