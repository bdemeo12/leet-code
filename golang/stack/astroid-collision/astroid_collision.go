package main

import "fmt"

// We are given an array asteroids of integers representing asteroids in a row.
//  The indices of the asteriod in the array represent their relative position in space.

// For each asteroid, the absolute value represents its size, and the sign represents
// its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

// Find out the state of the asteroids after all collisions. If two asteroids meet,
// the smaller one will explode. If both are the same size, both will explode.
// Two asteroids moving in the same direction will never meet.

func asteroidCollision(asteroids []int) []int {

	stack := []int{}

	for _, asteroid := range asteroids {

		for len(stack) > 0 || collision(stack[len(stack)-1], asteroid) {

			last := stack[len(stack)-1]
			// check if the incoming astroid is bigger than the last
			if last < asteroid {
				// remove last
				stack = stack[:len(stack)-1]
				continue
			} else if abs(last) == abs(asteroid) {
				stack = stack[:len(stack)-1]
			}

			asteroid = 0
			break

		}

		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func collision(a, b int) bool {

	if a > 0 && b < 0 {
		return true
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func larger(a, b int) int {
	if abs(a) > abs(b) {
		return a
	}

	return b
}

func main() {
	astroids := []int{10, 2, -5}
	fmt.Println(asteroidCollision(astroids))
}
