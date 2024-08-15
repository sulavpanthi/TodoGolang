package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func createFileIfNotExists(file_path string) {
	var file *os.File

	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		file, err = os.Create(file_path)
		if err != nil {
			fmt.Println("Cannot create a file", err)
			return
		}
		defer file.Close()
	}
}

func readTodos(file_path string) ([]Todo, error) {
	var (
		todos []Todo
		err   error
	)

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Cannot open file %v\n", err)
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Cannot read file %v\n", err)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Cannot convert string to integer %v\n", err)
		}

		isDone, err := strconv.ParseBool(record[2])
		if err != nil {
			fmt.Printf("Cannot convert string to boolean %v\n", err)
		}

		todos = append(todos, Todo{
			Id:     id,
			Title:  record[1],
			IsDone: isDone,
		})
	}
	return todos, err
}

func writeTodos(todos []Todo, file_path string) {
	file, err := os.Create(file_path)
	if err != nil {
		fmt.Printf("Cannot open file %v\n", err)
	}
	writer := csv.NewWriter(file)

	defer file.Close()
	defer writer.Flush()

	for _, record := range todos {
		idStr := fmt.Sprint(record.Id)
		isDoneStr := fmt.Sprint(record.IsDone)
		row := []string{idStr, record.Title, isDoneStr}
		err := writer.Write(row)
		if err != nil {
			fmt.Printf("Cannot write to file %v\n", err)
		}
	}

}

func addTodo() {
	createFileIfNotExists("todos.csv")

	var (
		title string
		todos []Todo
		err   error
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the title for your next todo: ")
	title, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("cannot read string properly", err)
	}
	title = title[:len(title)-1]

	todos, _ = readTodos("todos.csv")

	todos = append(todos, Todo{
		Id:     len(todos) + 1,
		Title:  title,
		IsDone: false,
	})

	writeTodos(todos, "todos.csv")
	fmt.Println("Todo has been successfully written...")
}

func listTodos() {
	var (
		todos []Todo
		err   error
	)

	createFileIfNotExists("todos.csv")

	todos, err = readTodos("todos.csv")
	if err != nil {
		fmt.Printf("Error occurred while reading file contents %v", err)
	}

	fmt.Print("\n****All Todos****\n\n")
	fmt.Print(todos, "\n\n")
}

func updateTodo() {
	var (
		todoId, updateChoice int
	)
	fmt.Print("Enter todo Id that you want to update: ")
	fmt.Scanln(&todoId)

	todos, err := readTodos("todos.csv") // Todo: Make readTodos dynamic to get todos in desired data structure type (e.g. slices, maps)
	if err != nil {
		fmt.Printf("Error while reading todos from csv file %v\n", err)
	}

	todo := todos[0]

	fmt.Println("Enter one of the choices below")
	fmt.Println("1. Update title")
	fmt.Println("2. Update completion status of todo")
	fmt.Println("3. Go back")

	fmt.Scanln(&updateChoice)

	switch updateChoice {
	case 1:
		var newTitle string
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter new title: ")
		newTitle, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Cannot read this input: ", err)
		}
		todo.Title = newTitle
	case 2:
		var (
			todoCompletionState string
		)
		fmt.Println("Is this todo completed? [Y/n]")
		fmt.Scanln(&todoCompletionState)

		switch todoCompletionState {
		case "Yes", "Y", "y", "yes":
			todo.IsDone = true
		case "No", "no", "N", "n":
			todo.IsDone = false
		default:
			fmt.Println("Todo has been marked complete. Please update again if this is not what you wanted.")
			todo.IsDone = true
		}
	case 3:
		return
	default:
		fmt.Println("Invalid option, try again later.")
		return
	}

	writeTodos(todos, "todos.csv")
}

func deleteTodo() {

}
