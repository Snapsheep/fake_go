package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task"
)

func main() {
	var action = os.Args[1]

	fmt.Printf("action : %s\n", action)

	switch action {
	case "add":
		{
			spName := strings.Split(os.Args[2], "=")
			spAge := strings.Split(os.Args[3], "=")

			name := spName[1]
			age, _ := strconv.Atoi(spAge[1])

			task.AddUsers(name, age)
		}
	case "list":
		task.ListUsers()
	default:
		fmt.Println("Invalid command : ", action)
	}

}
