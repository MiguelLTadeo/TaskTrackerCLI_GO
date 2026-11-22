/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MiguelLTadeo/TaskTrackerCLI_GO.git/internal"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid function!")
		return
	}

	funct := os.Args[1]

	switch funct {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Provide a description!")
			return
		}

		description := strings.Join(os.Args[2:], " ")

		err := internal.CreateTask(description)

		if err != nil {
			fmt.Println("Error creating task!")
			return
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Provide the ID and a new Description!")
			return
		}

		description := strings.Join(os.Args[3:], " ")

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID!")
			return
		}

		err = internal.UpdateTask(id, description)

		if err != nil {
			fmt.Println("Error updating task!")
			return
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Provide a ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID!")
			return
		}

		err = internal.DeleteTask(id)

		if err != nil {
			fmt.Println("Error deleting task!")
			return
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Provide a ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID!")
			return
		}
		err = internal.MarkInProgres(id)

		if err != nil {
			fmt.Println("Error marking task in progress!")
			return
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Provide a ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID!")
			return
		}
		err = internal.MarkDone(id)

		if err != nil {
			fmt.Println("Error marking task done!")
			return
		}
	case "list":
		if len(os.Args) < 3 {
			internal.ListAll()
			return
		}
		status := os.Args[2]

		internal.ListTasks(status)
	default:
		fmt.Println("Invalid function!")
		return
	}
}
