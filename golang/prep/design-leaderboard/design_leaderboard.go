package main

import (
	"fmt"
	"sort"
)

// Design a Leaderboard class, which has 3 functions:

// addScore(playerId, score): Update the leaderboard by adding score to the given player's score.
//  If there is no player with such id in the leaderboard, add him to the leaderboard with the given score.

// top(K): Return the score sum of the top K players.

// reset(playerId): Reset the score of the player with the given id to 0 (in other words erase it from the leaderboard).
//  It is guaranteed that the player was added to the leaderboard before calling this function.

// Initially, the leaderboard is empty.

type Leaderboard struct {
	PlayerScore map[int]int
}

func Constructor() Leaderboard {
	playerScore := make(map[int]int)
	return Leaderboard{
		PlayerScore: playerScore,
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	this.PlayerScore[playerId] += score
}

func (this *Leaderboard) Top(K int) int {

	scores := []int{}

	// get all scores
	for _, score := range this.PlayerScore {
		scores = append(scores, score)
	}

	sort.Slice(scores, func(a, b int) bool {
		return scores[a] > scores[b]
	})

	sum := 0
	for i := 0; i < K && i < len(scores); i++ {
		sum += scores[i]
	}

	return sum
}

func (this *Leaderboard) Reset(playerId int) {
	delete(this.PlayerScore, playerId)
}

func main() {

	obj := Constructor()

	obj.AddScore(1, 13)
	obj.AddScore(2, 93)
	obj.AddScore(3, 84)
	obj.AddScore(4, 6)
	obj.AddScore(5, 89)
	obj.AddScore(6, 31)
	obj.AddScore(7, 7)
	obj.AddScore(8, 1)
	obj.AddScore(9, 98)
	obj.AddScore(10, 42)

	fmt.Println(obj.Top(5)) // 406

	obj.Reset(1)
	obj.Reset(2)

	obj.AddScore(3, 76)
	obj.AddScore(4, 68)

	fmt.Println(obj.Top(1))

	obj.Reset(3)
	obj.Reset(4)

	obj.AddScore(2, 70)
	obj.Reset(2)
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
