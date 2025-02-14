package main

import "sort"

// There are n kids with candies. You are given an integer array candies,
// where each candies[i] represents the number of candies the ith kid has,
// and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if,
// after giving the ith kid all the extraCandies, they will have the greatest number of
// candies among all the kids, or false otherwise.

// Note that multiple kids can have the greatest number of candies.

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var output []bool

	// find kid with the most candies

	m := mostCandies(candies)

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= m+extraCandies {
			output = append(output, true)
			continue
		}
		output = append(output, false)
	}

	return output
}

func mostCandies(arr []int) int {
	var tmp = make([]int, len(arr))
	copy(tmp, arr)

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})

	return tmp[0]
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	result := kidsWithCandies(candies, extraCandies)
	for _, r := range result {
		println(r)
	}
}
