package main

import "fmt"

// You are given a list of preferences for n friends, where n is always even.

// For each person i, preferences[i] contains a list of friends sorted in the order of preference.
//  In other words, a friend earlier in the list is more preferred than a friend later in the list.
//  Friends in each list are denoted by integers from 0 to n-1.

// All the friends are divided into pairs. The pairings are given in a list pairs,
//  where pairs[i] = [xi, yi] denotes xi is paired with yi and yi is paired with xi.

// However, this pairing may cause some of the friends to be unhappy.
//  A friend x is unhappy if x is paired with y and there exists a friend u who is paired with v but:

// x prefers u over y, and
// u prefers x over v.

// Return the number of unhappy friends.

// hint: Create a matrix “rank” where rank[i][j] holds how highly friend ‘i' views ‘j’.
//  This allows for O(1) comparisons between people

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {

	// rank will be a map of a map. map person, to friend, to rank
	rankMap := make(map[int]map[int]int)
	for person := 0; person < len(preferences); person++ {
		rankMap[person] = make(map[int]int)

		for rank, friend := range preferences[person] {
			rankMap[person][friend] = rank // rankmap of person to friend to rank
		}
	}

	pairsMap := make(map[int]int)
	for i := 0; i < len(pairs); i++ {
		pairsMap[pairs[i][0]] = pairs[i][1]
		pairsMap[pairs[i][1]] = pairs[i][0]
	}

	unhappyFriends := make(map[int]bool)
	for x := 0; x < n; x++ {
		// if a pair perfers someone else, rank != 0. check if who they perfer perfers them higher than the person they are with
		y := pairsMap[x]

		// if x prefers someone else over y
		for _, u := range preferences[x] { // loop through anyone x perfers over their current partner
			if u == y { // are are at x partner y now
				break
			}

			// we found u -> someone x prefers of y
			v := pairsMap[u]                   // find u partner
			if rankMap[u][x] < rankMap[u][v] { // check to see if u prefers x over their current partner
				unhappyFriends[x] = true
			}

		}

	}

	return len(unhappyFriends)
}

func main() {
	n := 4
	preferences := [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}
	pairs := [][]int{{0, 1}, {2, 3}}
	fmt.Println(unhappyFriends(n, preferences, pairs)) // output: 2
}
