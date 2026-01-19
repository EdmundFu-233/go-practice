package main

import "fmt"

type Graph struct {
	nodes map[int][]int
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.nodes[u] = append(g.nodes[u], v)
	g.nodes[v] = append(g.nodes[v], u)
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)
		for _, neighbor := range g.nodes[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	fmt.Print("BFS from 1: ")
	g.BFS(1)
}
