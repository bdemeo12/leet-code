package main

import "fmt"

// In this problem, you are given an array of strings called wordsDict,
// which contains a list of words. You are also given two different strings word1 and word2
// which you can be certain are present in wordsDict. Your task is to find the shortest distance
// between word1 and word2 within wordsDict. The distance between two words is the number of words
// between them in the list (or the absolute difference between their indices in wordsDict).
// You need to consider the case where there may be multiple occurrences of word1 and/or word2,
// and you should find the minimum distance among all possible pairs of word1 and word2.

func shortestWordDistance(word1, word2 string, wordsDict []string) int {

	// loop through wordsDict to find word 1
	var minDistance int
	first := true
	var i, j = -1, -1

	for k := 0; k < len(wordsDict); k++ {
		if wordsDict[k] == word1 {
			i = k
		}

		if wordsDict[k] == word2 {
			j = k
		}

		// weve found both words
		if i != -1 && j != -1 {
			distance := j - i

			if first || minDistance > distance {
				minDistance = distance
				first = false
			}
		}
	}

	return minDistance
}

func main() {
	word1, word2 := "hello", "world"
	wordsDict := []string{"hello", "cat", "dog", "world", "test"}
	fmt.Println(shortestWordDistance(word1, word2, wordsDict))
}
