package main

import "fmt"

// You have a RecentCounter class which counts the number of recent requests within a certain time frame.

// Implement the RecentCounter class:

// RecentCounter() Initializes the counter with zero recent requests.

// int ping(int t) Adds a new request at time t, where t represents some time in milliseconds,
//  and returns the number of requests that has happened in the past 3000 milliseconds (including the new request).
//  Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].

// It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

type RecentCounter struct {
	start int

	records []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		start:   0,
		records: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	for this.start < len(this.records) && t-this.records[this.start] > 3000 {
		this.start++
	}

	this.records = append(this.records, t)

	return len(this.records) - this.start
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {

	counter := Constructor()
	fmt.Println(counter.Ping(1))    // returns 1
	fmt.Println(counter.Ping(100))  // returns 2
	fmt.Println(counter.Ping(3001)) // returns 3
	fmt.Println(counter.Ping(3002)) // returns 3
}
