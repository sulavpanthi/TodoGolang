# Todo Project in Go

## Overview

This project is a simple Todo application written in Go. It allows users to manage a list of todos, including adding, updating, and retrieving todo items. The project demonstrates basic usage of Go’s data structures and file handling.

## Features

- **Add Todo**: Add a new todo item with a title and completion status.
- **Update Todo**: Update the title and completion status of an existing todo item.
- **List Todos**: Retrieve and display all todo items.
- **Delete Todos**: Delete a particular todo.
- **Persistent Storage**: Save and load todos from a CSV file.

## Getting Started

### Prerequisites

- Go 1.18 or later installed on your machine.
- Basic knowledge of Go programming language.

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/sulavpanthi/TodoGolang.git
   ```

2. **Build the Project**

   ```bash
   cd cmd/cli
   go build -o main
   ```

### Usage

1. **Run the Application**

   ```bash
   cd cmd/cli
   ./main
   ```

### File Format

- **todos.csv**: The CSV file used for persistent storage. Each row represents a todo item with the following columns:
  - `ID`: Unique identifier for the todo item.
  - `Title`: The title of the todo item.
  - `IsDone`: Boolean value indicating whether the todo is completed.

### Project Structure

- `main.go`: The entry point of the application. Defines the Todo struct.
- `file_io.go`: Handles reading from and writing to the CSV file along with methods for manipulating todo items.

### Example

```go
package main

import (
	"fmt"
	"log"
)

// Example of creating a new Todo
func main() {
	todo := Todo{
		Id:     1,
		Title:  "Learn Go",
		IsDone: false,
	}
	fmt.Printf("Todo: %+v\n", todo)
}
```

### Contributing

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- Go programming language
- Go standard library for CSV handling
