package main

import (
	"container/heap"
	"fmt"
	"math"
)

type State struct {
	time, row, col, parity int
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTimeToReach(moveTime [][]int) int {
	DIRECTIONS := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	n, m := len(moveTime), len(moveTime[0])

	const INF = math.MaxInt32

	// dist[r][c][parity] stores earliest time we can reach (r,c)
	// with next move having parity (0 = next move 1 sec, 1 = next move 2 sec)
	dist := make([][][]int, n)
	for i := range dist {
		dist[i] = make([][]int, m)
		for j := range dist[i] {
			dist[i][j] = []int{INF, INF} // [parity 0, parity 1]
		}
	}

	pq := &MinHeap{}
	heap.Init(pq)

	// Start at (0,0), time=0, next move is 1 sec (parity = 0)
	heap.Push(pq, State{time: 0, row: 0, col: 0, parity: 0})
	dist[0][0][0] = 0

	for pq.Len() > 0 {
		state := heap.Pop(pq).(State)
		time, row, col, parity := state.time, state.row, state.col, state.parity

		// Skip if we already have a better time for this cell+parity
		if dist[row][col][parity] < time {
			continue
		}

		// Explore neighbors
		for _, dir := range DIRECTIONS {
			r, c := row+dir[0], col+dir[1]
			if r < 0 || r >= n || c < 0 || c >= m {
				continue
			}

			moveCost := 1
			if parity == 1 {
				moveCost = 2
			}

			arrivalTime := max(moveTime[r][c], time) + moveCost
			nextParity := 1 - parity // alternate move cost

			if arrivalTime < dist[r][c][nextParity] {
				dist[r][c][nextParity] = arrivalTime
				heap.Push(pq, State{
					time:   arrivalTime,
					row:    r,
					col:    c,
					parity: nextParity,
				})
			}
		}
	}

	// Return the best of reaching (n-1, m-1) with parity 0 or 1
	return min(dist[n-1][m-1][0], dist[n-1][m-1][1])
}

func main() {
	moveTime := [][]int{{0, 4}, {4, 4}}
	fmt.Println(minTimeToReach(moveTime)) // Should output 6
}
