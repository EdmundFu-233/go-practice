package algorithms

import (
	"fmt"
)

// TopologicalSort performs topological ordering of a directed acyclic graph
// using Kahn's algorithm (BFS-based). Returns the sorted nodes or an error if a cycle exists.
func TopologicalSort(nodes int, edges map[int][]int) ([]int, error) {
	inDegree := make([]int, nodes)
	for _, neighbors := range edges {
		for _, v := range neighbors {
			inDegree[v]++
		}
	}

	queue := make([]int, 0)
	for i := 0; i < nodes; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0, nodes)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)

		for _, v := range edges[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != nodes {
		return nil, fmt.Errorf("graph contains a cycle: %d/%d nodes processed", len(result), nodes)
	}

	return result, nil
}
