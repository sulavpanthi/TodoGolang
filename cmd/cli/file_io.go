package main

import (
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
	file, err := os.Open(file_path)
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

	fmt.Print("Enter the title for your next todo: ")
	fmt.Scan(&title)

	todos, _ = readTodos("todos.csv")

	file, err := os.Open("todos.csv")
	if err != nil {
		fmt.Printf("Cannot open file %v\n", err)
	}

	defer file.Close()

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
		file  *os.File
		err   error
	)

	// 1. create a new file if file does not exist
	createFileIfNotExists("todos.csv")

	// 2. Open and defer file close operation
	file, err = os.Open("todos.csv")
	if err != nil {
		fmt.Printf("Cannot open the file %v\n", err)
		return
	}

	defer file.Close()

	// 3. Create a new csv reader and read contents to a data structure
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Cannot read the file %v\n", err)
		}

		// convert string to int for id
		Id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Cannot convert string to integer %v\n", err)
		}

		// convert string to bool for isDone
		IsDone, err := strconv.ParseBool(record[2])
		if err != nil {
			fmt.Printf("Cannot convert string to boolean %v\n", err)
		}

		todos = append(todos, Todo{
			Id:     Id,
			Title:  record[1],
			IsDone: IsDone,
		})

	}
	fmt.Print("\n****All Todos****\n\n")
	fmt.Print(todos, "\n\n")

}

func updateTodo() {

}

func deleteTodo() {

}
