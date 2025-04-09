package main

import (
	"fmt"
	"sort"
)

// A company is planning to interview 2n people.
//  Given the array costs where costs[i] = [aCosti, bCosti],
//  the cost of flying the ith person to city a is aCosti, and the cost of flying the ith person to city b is bCosti.

// Return the minimum cost to fly every person to a city such that exactly n people arrive in each city.

// Greedy Algorithm
func twoCitySchedCost(costs [][]int) int {

	// sort by cost difference
	sort.Slice(costs, func(a, b int) bool {
		return costs[a][0]-costs[a][1] < costs[b][0]-costs[b][1]
	})

	// people with smallest cost differences, go to a
	// people with larger cost differences go to b

	n := len(costs) / 2
	totalCost := 0

	for i := 0; i < n; i++ {
		totalCost += costs[i][0]
	}

	for i := n; i < 2*n; i++ {
		totalCost += costs[i][1]
	}

	return totalCost
}

func main() {

	//costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}} //110
	costs := [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}} // 1859
	fmt.Println(twoCitySchedCost(costs))
}
