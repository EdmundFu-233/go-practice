package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, entry := range entries {
		info, _ := entry.Info()
		mode := "file"
		if entry.IsDir() {
			mode = "dir"
		}
		fmt.Printf("[%s] %s (%d bytes)\n", mode, entry.Name(), info.Size())
	}
}
