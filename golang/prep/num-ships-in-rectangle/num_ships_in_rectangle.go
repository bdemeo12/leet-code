package main

// Each ship is located at an integer point on the sea represented by a cartesian plane,
//  and each integer point may contain at most 1 ship.

// You have a function Sea.hasShips(topRight, bottomLeft)
//  which takes two points as arguments and returns true
// If there is at least one ship in the rectangle represented by the two points, including on the boundary.

// Given two points: the top right and bottom left corners of a rectangle,
//  return the number of ships present in that rectangle.
//  It is guaranteed that there are at most 10 ships in that rectangle.

// Submissions making more than 400 calls to hasShips will be judged Wrong Answer.
//  Also, any solutions that attempt to circumvent the judge will result in disqualification.

/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Sea struct {
 *     func hasShips(topRight, bottomLeft []int) bool {}
 * }
 */

// divide and conqure -> recurision

type Sea struct{}

// Define the method separately
func (s *Sea) hasShips(topRight, bottomLeft []int) bool {
	// Your logic here
	return true // or false based on logic
}

func countShips(sea Sea, topRight, bottomLeft []int) int {
	// we know we have traversed the whole grid when bottom exceeds top
	if bottomLeft[0] > topRight[0] || bottomLeft[1] > topRight[1] {
		return 0
	}

	if !sea.hasShips(topRight, bottomLeft) {
		return 0
	}

	// if a rectangle is a single point -> used to count
	if topRight[0] == bottomLeft[0] && topRight[1] == bottomLeft[1] {
		return 1
	}

	// find middle of x axis
	xMid := (topRight[0] + bottomLeft[0]) / 2

	// find middle of y axis
	yMid := (topRight[1] + bottomLeft[1]) / 2

	// divide into 4 sub grids
	return countShips(sea, []int{xMid, yMid}, bottomLeft) + countShips(sea, topRight, []int{xMid + 1, yMid + 1}) + countShips(sea, []int{xMid, topRight[1]}, []int{bottomLeft[0], yMid + 1}) + countShips(sea, []int{topRight[0], yMid}, []int{xMid + 1, bottomLeft[1]})
}

func main() {}
