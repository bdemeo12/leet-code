package main

import (
	"fmt"
	"strings"
)

// In the world of Dota2, there are two parties: the Radiant and the Dire.

// The Dota2 senate consists of senators coming from two parties.
//  Now the Senate wants to decide on a change in the Dota2 game.
//  The voting for this change is a round-based procedure.
//  In each round, each senator can exercise one of the two rights:

// Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
// Announce the victory: If this senator found the senators who still have rights to vote are all from the same party,
//  he can announce the victory and decide on the change in the game.
// Given a string senate representing each senator's party belonging.
//  The character 'R' and 'D' represent the Radiant party and the Dire party.
//  Then if there are n senators, the size of the given string will be n.

// The round-based procedure starts from the first senator to the last senator in the given order.
//  This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.

// Suppose every senator is smart enough and will play the best strategy for his own party.
//  Predict which party will finally announce the victory and change the Dota2 game. The output should be "Radiant" or "Dire".

func predictPartyVictory(senate string) string {

	queue := strings.Split(senate, "")

	rCount, dCount := 0, 0

	for _, q := range queue {
		if q == "R" {
			rCount++
		} else if q == "D" {
			dCount++
		}
	}

	if dCount == 0 {
		return "Radiant"
	} else if rCount == 0 {
		return "Dire"
	}

	bannedR, bannedD := 0, 0
	for dCount > 0 && rCount > 0 {
		s := queue[0]
		queue = queue[1:]

		if s == "R" {
			if bannedR == 0 {
				queue = append(queue, s)
				bannedD++
			} else {
				rCount--
				bannedR--
			}

		} else if s == "D" {
			if bannedD == 0 {
				queue = append(queue, s)
				bannedR++
			} else {
				bannedD--
				dCount--
			}
		}
	}

	// announce victory
	if queue[0] == "R" {
		return "Radiant"
	}

	return "Dire"
}

func banSenetorsRight(s string, queue []string) []string {

	for i, q := range queue {
		if q == s {
			queue[i] = "-"
			break
		}
	}

	return queue
}

func main() {
	//senate := "RDD" // output = dire
	//senate := "DDRRR" // DIRE
	senate := "RRDDD" // Radiant
	fmt.Println(predictPartyVictory(senate))
}
