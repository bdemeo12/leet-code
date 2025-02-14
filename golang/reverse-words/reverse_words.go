package main

// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters.
// The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words.
// The returned string should only have a single space separating the words. Do not include any extra spaces.

func reverseWords(s string) string {
	var output string
	runes := []rune(s)
	var tmp []rune

	for i := len(runes) - 1; i >= 0; i-- {
		if string(runes[i]) != " " {
			tmp = append([]rune{runes[i]}, tmp...)
			if i == 0 {
				return output + string(tmp)
			}
			continue
		}

		output = output + string(tmp) + string(runes[i])
		// clear tmp
		tmp = []rune{}
	}

	return ""
}

func main() {
	s := "  hello world  "
	result := reverseWords(s)
	println(result) // Output: "world hello
}

// func reverseWords(s string) string {
// 	var output string
// 	runes := []rune(s)
// 	var word1, word2 []rune
// 	i, j := 0, len(runes)-1

// 	for i < j {
// 		if string(runes[i]) != "" {
// 			word1 = append(word1, runes[i])
// 			i++
// 		}

// 		if string(runes[j]) != "" {
// 			word2 = append([]rune{runes[j]}, word2...)
// 			j++
// 		}

// 	}

// 	return output
// }
