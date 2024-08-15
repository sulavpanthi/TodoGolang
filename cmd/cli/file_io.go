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

}

func deleteTodo() {

}
