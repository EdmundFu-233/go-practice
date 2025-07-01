package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node int
	dist int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{} { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[:n-1]; return item }

func dijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0
	pq := &PriorityQueue{}
	heap.Push(pq, &Item{node: start, dist: 0})
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		if item.dist > dist[item.node] {
			continue
		}
		for _, edge := range graph[item.node] {
			newDist := dist[item.node] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(pq, &Item{node: edge.to, dist: newDist})
			}
		}
	}
	return dist
}

func main() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}
	dist := dijkstra(graph, 0)
	for node, d := range dist {
		fmt.Printf("Dist from 0 to %d: %d\n", node, d)
	}
}
