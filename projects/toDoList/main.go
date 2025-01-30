package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Struc of Todos
	type Todos struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}

	// How to use
	fmt.Println("-------How To Use-------")
	fmt.Print("[*]Creating a Task: 'create <write the task>'\n", "[*]To Show Existing Task: 'show'\n", "[*]Marking Done: 'done <id of the task>'\n", "[*]To Delete Task: 'remove <id of the task>'\n", "[*]To Close: 'close'\n", ">>>")

	//Counter
	c := 0

	//Create Tasks
	dx := make([]Todos, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {
		// Create Task
		case "create":
			dx = append(dx, Todos{
				Id:          c,
				Description: strings.Join(split[1:], " "),
				Status:      false,
			})
			c++
			fmt.Println(">>>Task Created")
		// Remove Task
		case "remove":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println(">>>Something Went Wrong!!")
			}
			for i, todo := range dx {
				if todo.Id == index {
					fmt.Println(">>>Task", todo.Description, "Removed")
					dx = slices.Delete(dx, i, i+1)
					break
				}
			}
		// Display The TODoList
		case "show":
			fmt.Println("-------TO DO LIST-------")
			for _, todo := range dx {
				flag := "[ ]"
				if todo.Status == true {
					flag = "[X]"
				}
				fmt.Println(flag, "ID:", todo.Id, todo.Description)
			}
		//Marking Done
		case "done":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println(">>>Something Went Wrong!!!")
			}
			for i, todo := range dx {
				if todo.Id == index {
					fmt.Println(">>>Task", todo.Description, "Done")
					dx[i].Status = true
				}
			}
		// CLosing The Program
		case "close":
			fmt.Println(">>>Good Bye!!!!")
			os.Exit(0)
		}
		fmt.Print(">>>")
	}
}
