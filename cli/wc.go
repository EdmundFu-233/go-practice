package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wc <filename>")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	lines, words, chars := 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		chars += len(line)
	}
	fmt.Printf("%d %d %d %s\n", lines, words, chars, os.Args[1])
}
