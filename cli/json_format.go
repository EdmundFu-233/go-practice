package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: json_format <json-string>")
		return
	}
	var raw interface{}
	if err := json.Unmarshal([]byte(os.Args[1]), &raw); err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}
	formatted, _ := json.MarshalIndent(raw, "", "  ")
	fmt.Println(string(formatted))
}
