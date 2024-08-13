package main

import "fmt"

type Todo struct {
	Id     int
	Title  string
	IsDone bool
}

func main() {
	for {
		fmt.Println("1. Add To-Do")
		fmt.Println("2. List To-Dos")
		fmt.Println("3. Update To-Do")
		fmt.Println("4. Remove To-Do")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTodo()
		case 2:
			listTodos()
		case 3:
			updateTodo()
		case 4:
			deleteTodo()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("This choice is invalid.")
		}
	}
}
