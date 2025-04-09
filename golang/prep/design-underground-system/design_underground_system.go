package main

import "fmt"

// An underground railway system is keeping track of customer travel times between different stations.
//  They are using this data to calculate the average time it takes to travel from one station to another.

// Implement the UndergroundSystem class:

// void checkIn(int id, string stationName, int t)
// A customer with a card ID equal to id, checks in at the station stationName at time t.
// A customer can only be checked into one place at a time.

// void checkOut(int id, string stationName, int t)
// A customer with a card ID equal to id, checks out from the station stationName at time t.

// double getAverageTime(string startStation, string endStation)

// Returns the average time it takes to travel from startStation to endStation.
// The average time is computed from all the previous traveling times from startStation to endStation
// that happened directly, meaning a check in at startStation followed by a check out from endStation.
// The time it takes to travel from startStation to endStation may be different from the time it takes
// to travel from endStation to startStation.
// There will be at least one customer that has traveled from startStation to endStation before getAverageTime is called.
// You may assume all calls to the checkIn and checkOut methods are consistent. If a customer checks in at time
// t1 then checks out at time t2, then t1 < t2. All events happen in chronological order.

type UndergroundSystem struct {
	Rides   map[int]Ride
	AvgTime map[string]Avg // map of startStationEndstation, values are different ride times
}

type Avg struct {
	totaltime int
	count     int
}

type Ride struct {
	id               int
	startStationName string
	EndStationName   string
	inTime           int
	outTime          int
}

func Constructor() UndergroundSystem {
	rides := make(map[int]Ride)
	avgTime := make(map[string]Avg)
	return UndergroundSystem{
		Rides:   rides,
		AvgTime: avgTime,
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {

	this.Rides[id] = Ride{
		id:               id,
		startStationName: stationName,
		inTime:           t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {

	ride := this.Rides[id]

	// now that a ride has ended, we should add it to our avg map
	key := fmt.Sprint(ride.startStationName + "_" + stationName)
	tripTime := t - ride.inTime

	avg := this.AvgTime[key]
	avg.totaltime += tripTime
	avg.count++

	this.AvgTime[key] = avg

	// delete from rides since tide is over
	delete(this.Rides, id)

}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	// get values from map for key
	key := fmt.Sprint(startStation + "_" + endStation)
	avg := this.AvgTime[key]

	return float64(avg.totaltime) / float64(avg.count)
}

func main() {

	obj := Constructor()
	obj.CheckIn(117077, "OZ13ZSIN", 42)
	obj.CheckOut(117077, "FU1D57HA", 125)
	obj.CheckIn(704903, "OZ13ZSIN", 211)
	obj.CheckIn(911239, "OZ13ZSIN", 245)
	obj.CheckOut(704903, "HG98HCFL", 340)
	obj.CheckIn(859828, "OZ13ZSIN", 355)

	fmt.Println(obj.GetAverageTime("OZ13ZSIN", "FU1D57HA"))

	obj.CheckOut(911239, "FU1D57HA", 383)
	fmt.Println(obj.GetAverageTime("OZ13ZSIN", "FU1D57HA"))
	fmt.Println(obj.GetAverageTime("OZ13ZSIN", "HG98HCFL"))

	// obj.CheckIn(45, "Leyton", 3)
	// obj.CheckIn(32, "Paradise", 8)
	// obj.CheckIn(27, "Leyton", 10)
	// obj.CheckOut(45, "Waterloo", 15)
	// obj.CheckOut(27, "Waterloo", 20)
	// obj.CheckOut(32, "Cambridge", 22)

	// fmt.Println(obj.GetAverageTime("Paradise", "Cambridge")) // 14
	// fmt.Println(obj.GetAverageTime("Leyton", "Waterloo"))    // 11

	// obj.CheckIn(10, "Leyton", 24)

	// fmt.Println(obj.GetAverageTime("Leyton", "Waterloo")) //11

	// obj.CheckOut(10, "Waterloo", 38)

	// fmt.Println(obj.GetAverageTime("Leyton", "Waterloo")) //12

}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
