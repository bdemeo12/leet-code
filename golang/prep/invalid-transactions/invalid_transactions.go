package main

import (
	"fmt"
	"strconv"
	"strings"
)

// A transaction is possibly invalid if:

// the amount exceeds $1000, or;
// if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.
// You are given an array of strings transaction where transactions[i] consists of comma-separated values representing
//  the name, time (in minutes), amount, and city of the transaction.

// Return a list of transactions that are possibly invalid. You may return the answer in any order.

type transaction struct {
	name    string
	time    int
	amount  int
	city    string
	str     string
	invalid bool
}

func invalidTransactions(transactions []string) []string {

	var trans []*transaction

	for _, t := range transactions {

		arr := strings.Split(t, ",")

		time, _ := strconv.Atoi(arr[1])
		amount, _ := strconv.Atoi(arr[2])

		trans = append(trans, &transaction{
			name:    arr[0],
			time:    time,
			amount:  amount,
			city:    arr[3],
			str:     t,
			invalid: false,
		})

	}

	invalid := []string{}

	for _, tran := range trans {
		if tran.amount > 1000 && !tran.invalid {
			invalid = append(invalid, tran.str)
			tran.invalid = true
			continue
		}
	}

	// brute force and check all other transactions against transaction
	for i := 0; i < len(trans); i++ {
		if trans[i].invalid { // its invalid lets skip
			continue
		}

		for j := 0; j < len(trans); j++ {
			if i == j { // dont compare with itself
				continue
			}

			tran := trans[i]
			tran2 := trans[j]

			// check if they are the same name in different cities
			if tran.name == tran2.name && tran.city != tran2.city {
				// check time difference
				if abs(tran.time, tran2.time) <= 60 {
					if !tran.invalid {
						invalid = append(invalid, tran.str)
						tran.invalid = true
					}

					if !tran2.invalid {
						invalid = append(invalid, tran2.str)
						tran2.invalid = true
					}
				}
			}
		}
	}

	return invalid
}

func abs(i, j int) int {
	if i < j {
		return j - i
	}

	return i - j
}

func main() {
	//t := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	//t := []string{"alex,676,260,bangkok", "bob,656,1366,bangkok", "alex,393,616,bangkok", "bob,820,990,amsterdam", "alex,596,1390,amsterdam"}
	t := []string{"bob,627,1973,amsterdam", "alex,387,885,bangkok", "alex,355,1029,barcelona", "alex,587,402,bangkok", "chalicefy,973,830,barcelona", "alex,932,86,bangkok", "bob,188,989,amsterdam"}
	// ["bob,627,1973,amsterdam",
	// "alex,387,885,bangkok",
	// "alex,355,1029,barcelona"]

	//[bob,627,1973,amsterdam
	//  alex,387,885,bangkok
	//  alex,355,1029,barcelona
	//  alex,355,1029,barcelona]
	fmt.Println(invalidTransactions(t))
}
