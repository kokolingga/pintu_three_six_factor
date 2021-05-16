package main

import "fmt"

func checkIfHasSixFactors(inputNumber int) bool {
	lastValidTracking := 0

	// expected lastValidTracking will be 4 (6-2)
	// no need to check inputNumber % 1
	// no need to check inputNumber % inputNumber
	for i := 2; i < inputNumber; i++ {
		if inputNumber%i == 0 {
			lastValidTracking++

			if lastValidTracking > 4 {
				return false
			}
		}
	}

	if lastValidTracking == 4 {
		return true
	} else {
		return false
	}
}

func getTotalNumbersWithSixFactors(maxRange int) int {

	statusTracker := make(map[int]string, maxRange)
	var totalNumbersWithSixFactors int

	// number below 6 most likely have less than 6 factor.
	// so we'll start from 6.
	for i := 6; i <= maxRange; i++ {
		if statusTracker[i] == "" {
			if checkIfHasSixFactors(i) == true {
				markedTheSiblings(i, maxRange, statusTracker)
				totalNumbersWithSixFactors += 1
			}
		}
	}

	return totalNumbersWithSixFactors
}

func markedTheSiblings(numberWithSixFactor int, maxRange int, statusTracker map[int]string) {
	counter := 2

	for {
		if numberWithSixFactor*counter > maxRange {
			break
		}

		message := fmt.Sprintf("NOT-A-SIX-FACTOR (SIBLING OF %v)", numberWithSixFactor)
		// numberWithSixFactor already has 6 factors.
		// after the multiplication, the "bigger" siblings most likely have more than 6 factors.
		statusTracker[numberWithSixFactor*counter] = message

		counter++
	}
}

func main() {
	maxRangeToCheck := 262144 // H(262144) = 13208 (took 1m 39s)

	totalNumbersWithSixFactors := getTotalNumbersWithSixFactors(maxRangeToCheck)

	fmt.Println(totalNumbersWithSixFactors)

	/*
		// this will take up to (134217728/262144) * 1m 39s => 14+ hours.

		maxRangeToCheck := 134217728

		totalNumbersWithSixFactors := getTotalNumbersWithSixFactors(maxRangeToCheck)
		fmt.Println(totalNumbersWithSixFactors)
	*/
}
