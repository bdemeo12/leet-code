package main

func largestValue(m map[string]int) int {

	var first bool
	var largest int

	for _, val := range m {

		if first || val > largest {
			largest = val
		}
	}

	return largest
}
