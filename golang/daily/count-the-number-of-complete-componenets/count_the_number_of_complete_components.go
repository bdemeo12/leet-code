package main

import "fmt"

// You are given an integer n. There is an undirected graph with n vertices,
// numbered from 0 to n - 1. You are given a 2D integer array edges where edges[i] = [ai, bi]
// denotes that there exists an undirected edge connecting vertices ai and bi.

// Return the number of complete connected components of the graph.

// A connected component is a subgraph of a graph in which there exists a path between any two vertices,
// and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

// A connected component is said to be complete if there exists an edge between every pair of its vertices.

func countCompleteComponents(n int, edges [][]int) int {

	// create a graph
	graph := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		a := edges[i][0]
		b := edges[i][1]

		//set in map
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make(map[int]bool)
	count := 0

	// for each connected component, count the number of nodes and edges in each component
	for i := 0; i < n; i++ {
		if _, ok := visited[i]; !ok {
			component := []int{}
			edgeCount := dfs(graph, i, &component, visited)
			nodesCount := len(component)

			if edgeCount/2 == nodesCount*(nodesCount-1)/2 {
				count++
			}
		}
	}

	return count
}

func dfs(graph map[int][]int, node int, component *[]int, visited map[int]bool) int {
	*component = append(*component, node)
	visited[node] = true

	edges := 0
	for _, neighbor := range graph[node] {
		edges++
		if !visited[neighbor] {
			edges += dfs(graph, neighbor, component, visited)
		}
	}

	return edges
}

func main() {
	// n := 6
	// edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}
	// output = 3

	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}
	fmt.Println(countCompleteComponents(n, edges))
}
