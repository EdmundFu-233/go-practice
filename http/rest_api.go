package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items = []Item{{1, "item1"}, {2, "item2"}}
	mu    sync.Mutex
	nextID = 3
)

func getItems(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/items", getItems)
	fmt.Println("REST API on :8080")
	http.ListenAndServe(":8080", nil)
}
