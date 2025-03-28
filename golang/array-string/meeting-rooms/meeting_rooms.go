package main

import (
	"fmt"
	"sort"
)

// The problem presents a scenario in which an individual has multiple meetings scheduled,
//  represented by time intervals. Each interval is an array of two values [start_i, end_i],
//  where start_i is the time a particular meeting begins,
//  and end_i is the time the meeting ends.
//  The intervals array contains all such meeting time intervals.

// The task is to determine whether the person can attend all the scheduled meetings without any overlaps.
//  More specifically, no two meetings can occur at the same time.
//  The person can only attend all meetings if, for any two meetings,
//  one finishes before the other starts.

func canAttendMeetings(intervals [][]int) bool {

	// sort slice by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	// if the start time of the current meeting is before the end of the pervious meeting, return false
	for i := 1; i < len(intervals); i++ {
		currentStartTime := intervals[i][0]
		previousEndTime := intervals[i-1][1]

		if currentStartTime < previousEndTime {
			return false
		}
	}

	return true
}

func main() {

	input := [][]int{{0, 30}, {5, 10}, {15, 20}} //Output:= false  Overlapping meetings

	// Input: [[7,10],[2,4]]
	// Output: true   // No overlaps
	fmt.Println(canAttendMeetings(input))
}
