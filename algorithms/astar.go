// Package algorithms provides common algorithm implementations.
package algorithms

import (
	"container/heap"
	"math"
)

// AStar implements the A* pathfinding algorithm.
// It finds the shortest path from start to goal on a grid.
// 0 = walkable, 1 = obstacle.
type AStar struct {
	grid  [][]int
	rows  int
	cols  int
}

// Node represents a point on the grid with A* metadata.
type Node struct {
	X, Y     int
	G, H, F  float64
	Parent   *Node
	Index    int // heap index
}

// PriorityQueue implements heap.Interface for A* nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.Index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.Index = -1
	*pq = old[:n-1]
	return node
}

// NewAStar creates a new A* solver for the given grid.
func NewAStar(grid [][]int) *AStar {
	return &AStar{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

// heuristic calculates Manhattan distance.
func (a *AStar) heuristic(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
}

// neighbors returns valid adjacent walkable cells.
func (a *AStar) neighbors(x, y int) [][2]int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var result [][2]int
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < a.rows && ny >= 0 && ny < a.cols && a.grid[nx][ny] == 0 {
			result = append(result, [2]int{nx, ny})
		}
	}
	return result
}

// FindPath finds the shortest path from (sx,sy) to (gx,gy).
// Returns the path as a slice of coordinates (inclusive), or nil if no path exists.
func (a *AStar) FindPath(sx, sy, gx, gy int) [][2]int {
	start := &Node{X: sx, Y: sy, G: 0}
	start.H = a.heuristic(sx, sy, gx, gy)
	start.F = start.G + start.H

	pq := &PriorityQueue{start}
	heap.Init(pq)

	visited := make(map[[2]int]bool)
	gScore := make(map[[2]int]float64)
	gScore[[2]int{sx, sy}] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		key := [2]int{current.X, current.Y}

		if current.X == gx && current.Y == gy {
			// Reconstruct path
			var path [][2]int
			for n := current; n != nil; n = n.Parent {
				path = append([][2]int{{n.X, n.Y}}, path...)
			}
			return path
		}

		if visited[key] {
			continue
		}
		visited[key] = true

		for _, nb := range a.neighbors(current.X, current.Y) {
			nk := [2]int{nb[0], nb[1]}
			if visited[nk] {
				continue
			}

			tentativeG := gScore[key] + 1
			if oldG, ok := gScore[nk]; !ok || tentativeG < oldG {
				gScore[nk] = tentativeG
				neighbor := &Node{
					X:      nb[0],
					Y:      nb[1],
					G:      tentativeG,
					H:      a.heuristic(nb[0], nb[1], gx, gy),
					Parent: current,
				}
				neighbor.F = neighbor.G + neighbor.H
				heap.Push(pq, neighbor)
			}
		}
	}

	return nil // No path found
}
