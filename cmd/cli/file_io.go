package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type DataStructure interface {
	Add(Todo)
	GetTodos() interface{}
}

type SliceDS struct {
	data []Todo
}

func (slice *SliceDS) Add(todo Todo) {
	slice.data = append(slice.data, todo)
}

func (slice *SliceDS) GetTodos() interface{} {
	return slice.data
}

type MapDS struct {
	data map[int]Todo
}

func (mapDS *MapDS) Add(todo Todo) {
	mapDS.data[todo.Id] = todo
}

func (mapDS *MapDS) GetTodos() interface{} {
	return mapDS.data
}

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

func readTodos(file_path string, dataStructure DataStructure) error {
	var (
		err error
	)

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Cannot open file %v\n", err)
		return err
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

		// todos = append(todos, Todo{
		// 	Id:     id,
		// 	Title:  record[1],
		// 	IsDone: isDone,
		// })
		dataStructure.Add(Todo{
			Id:     id,
			Title:  record[1],
			IsDone: isDone,
		})
	}
	return err
}

func writeTodos(todoDataStructure DataStructure, file_path string) {
	file, err := os.Create(file_path)
	if err != nil {
		fmt.Printf("Cannot open file %v\n", err)
	}
	writer := csv.NewWriter(file)

	defer file.Close()
	defer writer.Flush()

	switch todos := todoDataStructure.GetTodos().(type) {
	case []Todo:
		for _, record := range todos {
			idStr := fmt.Sprint(record.Id)
			isDoneStr := fmt.Sprint(record.IsDone)
			row := []string{idStr, record.Title, isDoneStr}
			err := writer.Write(row)
			if err != nil {
				fmt.Printf("Cannot write to file %v\n", err)
			}
		}
	case map[int]Todo:
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

}

func addTodo() {
	createFileIfNotExists("todos.csv")

	var (
		title string
		err   error
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the title for your next todo: ")
	title, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("cannot read string properly", err)
	}
	title = title[:len(title)-1]

	todoDataStructure := SliceDS{}
	err = readTodos("todos.csv", &todoDataStructure)
	if err != nil {
		fmt.Printf("Cannot read todos %v", err)
	}

	todoDataStructure.Add(Todo{
		Id:     len(todoDataStructure.data) + 1,
		Title:  title,
		IsDone: false,
	})

	writeTodos(&todoDataStructure, "todos.csv")
	fmt.Println("Todo has been successfully written...")
}

func listTodos() {
	var (
		err error
	)

	createFileIfNotExists("todos.csv")

	todoDataStructure := SliceDS{}
	err = readTodos("todos.csv", &todoDataStructure)
	if err != nil {
		fmt.Printf("Error occurred while reading file contents %v", err)
	}

	fmt.Print("\n****All Todos****\n\n")
	fmt.Print(todoDataStructure.data, "\n\n")
}

func updateTodo() {
	var (
		todoId, updateChoice int
	)
	fmt.Print("Enter todo Id that you want to update: ")
	fmt.Scanln(&todoId)

	todoDataStructure := SliceDS{}
	err := readTodos("todos.csv", &todoDataStructure) // Todo: Make readTodos dynamic to get todos in desired data structure type (e.g. slices, maps)
	if err != nil {
		fmt.Printf("Error while reading todos from csv file %v\n", err)
	}

	todo := todoDataStructure.data[0]

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

	writeTodos(&todoDataStructure, "todos.csv")
}

func deleteTodo() {

}
