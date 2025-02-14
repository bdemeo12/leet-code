package main

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a new string that is formed from the original string by deleting some
// (can be none) of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

func isSubsequence(s string, t string) bool {

	if s == "" {
		return true
	}

	// convert strings to arrs
	sArr := []rune(s)
	tArr := []rune(t)
	i, j := 0, 0

	for i < len(sArr) && j < len(tArr) {
		if string(sArr[i]) == string(tArr[j]) {
			i++
			j++
		} else {
			j++
		}

		if i == len(sArr) {
			return true
		}
	}

	return false
}

// // convert strings to arrs
// sArr := []rune(s)
// tArr := []rune(t)
// var count int

// for i := 0; i < len(tArr); i++ {
// 	for j := 0; j < len(sArr); j++ {
// 		if string(tArr[i]) == string(sArr[j]) {
// 			count++
// 			// break into the outer loop because we found a char
// 			break
// 		}
// 	}
// }

// return count == len(sArr)
