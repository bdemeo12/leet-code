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

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	result := largestValue(m)
	println(result) // Output: 3
}
