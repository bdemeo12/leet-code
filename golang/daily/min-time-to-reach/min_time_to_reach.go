package main

import (
	"container/heap"
	"fmt"
)

type State struct {
	time, row, col int
}

type MinHeap []State

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	DIRECTIONS := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	n, m := len(moveTime), len(moveTime[0])

	// Earliest time each cell is reached
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// Min-heap by time: (time, row, col)
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, State{time: 0, row: 0, col: 0})
	dist[0][0] = 0

	for pq.Len() > 0 {
		state := heap.Pop(pq).(State)
		time, row, col := state.time, state.row, state.col

		// Skip if already visited with a better time
		if dist[row][col] < time {
			continue
		}

		// Explore 4 neighbors
		for _, dir := range DIRECTIONS {
			r, c := row+dir[0], col+dir[1]
			if r < 0 || r >= n || c < 0 || c >= m {
				continue
			}

			// Time to enter (r, c) is max(wait, now) + 1
			nextTime := max(moveTime[r][c], time) + 1

			// Update if this path is faster
			if dist[r][c] == -1 || nextTime < dist[r][c] {
				dist[r][c] = nextTime
				heap.Push(pq, State{time: nextTime, row: r, col: c})
			}
		}
	}

	return dist[n-1][m-1]
}

// type Point struct {
// 	row, col int
// }

// type State struct {
// 	point Point
// 	time  int
// }

// func minTimeToReach(moveTime [][]int) int {
// 	n := len(moveTime)
// 	m := len(moveTime[0])

// 	// Directions → up, down, left, right
// 	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 	// Track visited cells
// 	visited := make([][]bool, n)
// 	for i := range visited {
// 		visited[i] = make([]bool, m)
// 	}

// 	// BFS queue → start from (0,0) at time=0
// 	queue := []State{{point: Point{row: 0, col: 0}, time: 0}}
// 	visited[0][0] = true

// 	for len(queue) > 0 {
// 		curr := queue[0]
// 		queue = queue[1:]

// 		// If we reached the target
// 		if curr.point.row == n-1 && curr.point.col == m-1 {
// 			return curr.time
// 		}

// 		// Explore all neighbors
// 		for _, dir := range directions {
// 			newRow := curr.point.row + dir.row
// 			newCol := curr.point.col + dir.col

// 			// Bounds + not visited
// 			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < m && !visited[newRow][newCol] {
// 				visited[newRow][newCol] = true

// 				nextMoveTime := moveTime[newRow][newCol]
// 				arrivalTime := curr.time + 1 // takes 1 second to move

// 				// Wait if arrivalTime < room's available time
// 				if arrivalTime < nextMoveTime {
// 					arrivalTime = nextMoveTime
// 				}

// 				queue = append(queue, State{
// 					point: Point{row: newRow, col: newCol},
// 					time:  arrivalTime,
// 				})
// 			}
// 		}
// 	}

// 	return -1 // Should not reach here
// }

func main() {
	moveTime := [][]int{{0, 4}, {4, 4}}
	fmt.Println(minTimeToReach(moveTime))
}
