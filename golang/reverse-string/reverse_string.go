package main

func reverseString(s string) string {

	i, j := 0, len(s)-1

	sArr := []rune(s)

	for i != j {

		sArr[i], sArr[j] = sArr[j], sArr[i]

		i++
		j--
	}

	return string(sArr)
}

func main() {
	s := "hello"
	result := reverseString(s)
	println(result) // Output: "olleh"
}
