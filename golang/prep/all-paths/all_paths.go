package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	target := len(graph) - 1
	paths := [][]int{}

	// queue to hold paths to explore
	queue := [][]int{{0}}

	for len(queue) > 0 {
		// pop first path from queue
		path := queue[0]
		queue = queue[1:]

		// get the last node in the path
		lastNode := path[len(path)-1]

		if lastNode == target {
			paths = append(paths, path)
		}

		// otherwise explore neighbors
		for _, neighbor := range graph[lastNode] {
			// make a copy of path before adding neighbor
			newPath := append([]int{}, path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, newPath)
		}
	}

	return paths
}

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
