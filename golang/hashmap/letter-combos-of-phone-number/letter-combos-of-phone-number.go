package main

import "fmt"

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combos := []string{""}

	for _, digit := range digits {
		newCombos := []string{}
		for _, combo := range combos {
			for _, letter := range phoneMap[digit-'2'] {
				newCombos = append(newCombos, combo+string(letter))
			}
		}
		combos = newCombos
	}

	return combos
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
