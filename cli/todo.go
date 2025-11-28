package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	todos := []Todo{
		{1, "Learn Go", false},
		{2, "Build CLI", false},
		{3, "Write tests", true},
	}
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile("todos.json", data, 0644)
	fmt.Println("Todos saved to todos.json")
}
