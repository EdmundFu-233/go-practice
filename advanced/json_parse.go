package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Tags  []string `json:"tags,omitempty"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}

func main() {
	jsonStr := `{"name":"Zhihao","age":22,"tags":["go","dev"],"meta":{"city":"Waterloo"}}`
	var user User
	json.Unmarshal([]byte(jsonStr), &user)
	fmt.Printf("%+v\n", user)

	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))
}
