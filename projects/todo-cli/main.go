package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Priority    string     `json:"priority"`
}

type TodoList struct {
	Todos    []Todo `json:"todos"`
	NextID   int    `json:"next_id"`
	filename string `json:"-"`
}

func NewTodoStore(filename string) *TodoList {
	store := &TodoList{
		Todos:    []Todo{},
		NextID:   1,
		filename: filename,
	}
	err := store.load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todo list: %v\n", err)
		os.Exit(1)
	}
	return store
}
func (s *TodoList) load() error {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, s)
}

func (s *TodoList) save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filename, data, 0644)
}

func (s *TodoList) add(Title, Description, Priority string) error {
	if Priority == "" {
		Priority = "medium"
	}
	todo := Todo{
		ID:          s.NextID,
		Title:       Title,
		Description: Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		Priority:    Priority,
	}
	s.Todos = append(s.Todos, todo)
	s.NextID++
	return s.save()
}

func (s *TodoList) list() {
	if len(s.Todos) == 0 {
		println("No todos found.")
		return
	}
	for _, todo := range s.Todos {
		status := "⬜"
		if todo.Completed {
			status = "✅"
		}

		priority := ""
		switch todo.Priority {
		case "high":
			priority = "🔴"
		case "medium":
			priority = "🟡"
		case "low":
			priority = "🟢"
		}

		fmt.Printf("%s [%d] %s %s\n", status, todo.ID, priority, todo.Title)
		if todo.Description != "" {
			fmt.Printf("    %s\n", todo.Description)
		}
		fmt.Printf("    Created: %s\n", todo.CreatedAt.Format("2006-01-02 15:04"))
		if todo.Completed && todo.CompletedAt != nil {
			fmt.Printf("    Completed: %s\n", todo.CompletedAt.Format("2006-01-02 15:04"))
		}
		fmt.Println()
	}

}

func (s *TodoList) complete(id int) error {
	for i, todo := range s.Todos {
		if todo.ID == id {
			if todo.Completed {
				return fmt.Errorf("todo %d is already completed", id)
			}
			now := time.Now()
			s.Todos[i].Completed = true
			s.Todos[i].CompletedAt = &now
			return s.save()
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (s *TodoList) uncomplete(id int) error {
	for i, todo := range s.Todos {
		if todo.ID == id {
			if !todo.Completed {
				return fmt.Errorf("todo %d is not completed", id)
			}
			s.Todos[i].Completed = false
			s.Todos[i].CompletedAt = nil
			return s.save()
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (s *TodoList) delete(id int) error {
	for i, todo := range s.Todos {
		if todo.ID == id {
			s.Todos = append(s.Todos[:i], s.Todos[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (s *TodoList) update(id int, Title, Description, Priority string) error {
	if Priority == "" {
		Priority = "medium"
	}
	for i, todo := range s.Todos {
		if todo.ID == id {
			s.Todos[i].Title = Title
			s.Todos[i].Description = Description
			s.Todos[i].Priority = Priority
			return s.save()
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (s *TodoList) search(query string) {
	found := false
	for _, todo := range s.Todos {
		if strings.Contains(todo.Title, query) || strings.Contains(todo.Description, query) {
			status := "⬜"
			if todo.Completed {
				status = "✅"
			}

			priority := ""
			switch todo.Priority {
			case "high":
				priority = "🔴"
			case "medium":
				priority = "🟡"
			case "low":
				priority = "🟢"
			}

			fmt.Printf("%s [%d] %s %s\n", status, todo.ID, priority, todo.Title)
			if todo.Description != "" {
				fmt.Printf("    %s\n", todo.Description)
			}
			fmt.Printf("    Created: %s\n", todo.CreatedAt.Format("2006-01-02 15:04"))
			if todo.Completed && todo.CompletedAt != nil {
				fmt.Printf("    Completed: %s\n", todo.CompletedAt.Format("2006-01-02 15:04"))
			}
			fmt.Println()
			found = true
		}
	}
	if !found {
		println("No matching todos found.")
	}
}

func printUsage() {
	fmt.Print(`
📋 Todo CLI - Usage:

  add <title> [-d description] [-p priority]     Add a new todo
  list                                      List todos (-a for all including completed)
  complete <id>                                  Mark todo as completed
  uncomplete <id>                                Mark todo as incomplete
  delete <id>                                    Delete a todo
  update <id> [-t title] [-d description] [-p priority]  Update a todo
  search <query>                                 Search todos
  help                                           Show this help message

Priority levels: low, medium, high (default: medium)

Examples:
  todo add "Buy groceries" -d "Milk, eggs, bread" -p high
  todo list
  todo complete 1
  todo update 2 -t "New title" -p low
  todo search "groceries"
`)
}

func main() {
	store := NewTodoStore("todos.json")

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Title is required for adding a todo.")
			return
		}
		title := os.Args[2]
		description := ""
		priority := "medium"
		for i := 3; i < len(os.Args); i++ {
			if os.Args[i] == "-d" && i+1 < len(os.Args) {
				description = os.Args[i+1]
				i++
			} else if os.Args[i] == "-p" && i+1 < len(os.Args) {
				priority = os.Args[i+1]
				i++
			}
		}
		err := store.add(title, description, priority)
		if err != nil {
			fmt.Printf("Error adding todo: %v\n", err)
		} else {
			fmt.Println("Todo added successfully.")
		}
	case "list":
		store.list()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required to complete a todo.")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		err := store.complete(id)
		if err != nil {
			fmt.Printf("Error completing todo: %v\n", err)
		} else {
			fmt.Println("Todo marked as completed.")
		}
	case "uncomplete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required to uncomplete a todo.")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		err := store.uncomplete(id)
		if err != nil {
			fmt.Printf("Error uncompleting todo: %v\n", err)
		} else {
			fmt.Println("Todo marked as incomplete.")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required to delete a todo.")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		err := store.delete(id)
		if err != nil {
			fmt.Printf("Error deleting todo: %v\n", err)
		} else {
			fmt.Println("Todo deleted successfully.")
		}
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required to update a todo.")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		title := ""
		description := ""
		priority := ""
		for i := 3; i < len(os.Args); i++ {
			if os.Args[i] == "-t" && i+1 < len(os.Args) {
				title = os.Args[i+1]
				i++
			} else if os.Args[i] == "-d" && i+1 < len(os.Args) {
				description = os.Args[i+1]
				i++
			} else if os.Args[i] == "-p" && i+1 < len(os.Args) {
				priority = os.Args[i+1]
				i++
			}
		}
		err := store.update(id, title, description, priority)
		if err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
		} else {
			fmt.Println("Todo updated successfully.")
		}
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Error: Query is required to search todos.")
			return
		}
		query := os.Args[2]
		store.search(query)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}
