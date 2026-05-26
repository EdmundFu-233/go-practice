// Package algorithms implements classical algorithms in Go.
//
// Floyd-Warshall: All-Pairs Shortest Path
//
// Finds shortest paths between all pairs of vertices in a weighted
// directed graph. Handles negative edge weights (but no negative cycles).
//
// Complexity: O(V³) time, O(V²) space.
package algorithms

import "math"

const inf = math.MaxInt32 >> 1 // avoid overflow on addition

// FloydWarshall computes the all-pairs shortest path distances for a graph
// represented as an adjacency matrix. dist[i][j] is the weight of edge i→j,
// or inf if no edge exists.
//
// Returns the shortest-path distance matrix and a next-hop matrix for
// path reconstruction. A negative cycle is signaled by a negative value
// on the diagonal.
func FloydWarshall(dist [][]int) (shortest [][]int, next [][]int, hasNegCycle bool) {
	n := len(dist)

	// Initialize
	shortest = make([][]int, n)
	next = make([][]int, n)
	for i := 0; i < n; i++ {
		shortest[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			shortest[i][j] = dist[i][j]
			if dist[i][j] != inf && i != j {
				next[i][j] = j
			} else {
				next[i][j] = -1
			}
		}
	}

	// Floyd-Warshall core: try every vertex as intermediate
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if shortest[i][k] != inf && shortest[k][j] != inf {
					candidate := shortest[i][k] + shortest[k][j]
					if candidate < shortest[i][j] {
						shortest[i][j] = candidate
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}

	// Check for negative cycles
	for i := 0; i < n; i++ {
		if shortest[i][i] < 0 {
			hasNegCycle = true
			return
		}
	}

	return
}

// ReconstructPath rebuilds the shortest path from u to v using the next-hop matrix.
func ReconstructPath(next [][]int, u, v int) []int {
	if next[u][v] == -1 {
		return nil // no path
	}
	path := []int{u}
	for u != v {
		u = next[u][v]
		path = append(path, u)
	}
	return path
}
