package main

// You have a long flowerbed in which some of the plots are planted,
// and some are not. However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's,
// where 0 means empty and 1 means not empty, and an integer n,
// return true if n new flowers can be planted in the flowerbed without violating
// the no-adjacent-flowers rule and false otherwise.

func canPlaceFlowers(flowerbed []int, n int) bool {
	// because flowers cannot be adjacent, we can only plant a flower every other
	if n > len(flowerbed)/2 {
		return true
	}

	for i := 1; i < len(flowerbed); i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--

			// if n == 0 {
			// 	break
			// }
		}
	}

	return n == 0
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1

	result := canPlaceFlowers(flowerbed, n)
	println(result)
}
