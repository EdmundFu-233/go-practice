package main

import (
	"fmt"
	"os"
)

func main() {
	content := []byte("Hello, Go!\n")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File written successfully")
}
