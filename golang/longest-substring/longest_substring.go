package main

// Given a string s and an integer k, return the length of the longest substring of s such
// that the frequency of each character in this substring is greater than or equal to k.

// if no such substring exists, return 0.

// sliding window

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	return 0
}

func main() {
	s := "aaabb"
	k := 3
	result := longestSubstring(s, k)
	println(result) // Output: 3
}

//sArr := []rune(s)
// counter := 1
// var tracker = make(map[rune]int)

// for i := 1; i < len(sArr); i++ {
// 	if sArr[i-1] == sArr[i] {
// 		counter++
// 		continue
// 	}

// 	val, ok := tracker[sArr[i-1]]
// 	if !ok || counter > val {
// 		tracker[sArr[i-1]] = counter
// 		counter = 0
// 	}
// }

// maxVal := 0
// first := true

// for _, value := range tracker {
// 	if first || (value > maxVal && value >= k) {
// 		maxVal = value
// 		first = false
// 	}
// }

// return maxVal
