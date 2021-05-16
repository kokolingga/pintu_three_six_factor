package main

import "fmt"

func getFactors(inputNumber int) []int {
	factors := make([]int, 0)

	for i := 1; i <= inputNumber; i++ {
		if inputNumber%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func checkIfHasSixFactors(inputNumber int) bool {
	lastValidTracking := 0

	for i := 1; i <= inputNumber; i++ {
		if inputNumber%i == 0 {
			lastValidTracking++

			if lastValidTracking > 6 {
				//fmt.Printf("%v is greater than 6\n", lastValidTracking)
				return false
			}
		}
	}

	if lastValidTracking == 6 {
		return true
	} else {
		//fmt.Printf("%v is less than 6\n", lastValidTracking)
		return false
	}
}

func getNumberOfFactors(inputNumber int) int {
	return len(getFactors(inputNumber))
}

// func getNumbersWithSixFactors(maxRange int) []int {
func getNumbersWithSixFactors(maxRange int) int {

	statusTracker := make(map[int]string, maxRange)
	// var numbersWithSixFactors []int
	var numbersWithSixFactorsAmount int

	// number below 6 most likely have less than 6 factor.
	// so we'll start from 6.
	for i := 6; i <= maxRange; i++ {
		if statusTracker[i] == "" {
			// numbers "i" has not been checked
			// fmt.Printf("\n%v is empty\n", i)

			// check how many factors the number has

			// (wrong approach. too long)
			// numberOfFactors := getNumberOfFactors(i)
			// fmt.Printf("number of factors : %v\n", numberOfFactors)

			// if numberOfFactors == 6 {
			if checkIfHasSixFactors(i) == true {
				// fmt.Printf("=== %v is 6 factor ===\n", i)
				// numbersWithSixFactors = append(numbersWithSixFactors, i)
				markedTheSiblings(i, maxRange, statusTracker)
				numbersWithSixFactorsAmount += 1
			}
		} else {
			// numbers "i" has been mark by "markedTheSiblings" function
			// fmt.Printf("\n%v : % v\n", i, statusTracker[i])

		}
	}

	// return numbersWithSixFactors
	return numbersWithSixFactorsAmount
}

func markedTheSiblings(numberWithSixFactor int, maxRange int, statusTracker map[int]string) {
	counter := 2

	for {
		if numberWithSixFactor*counter > maxRange {
			break
		}

		message := fmt.Sprintf("NOT-A-SIX-FACTOR (SIBLING OF %v)", numberWithSixFactor)
		statusTracker[numberWithSixFactor*counter] = message

		counter++
	}
}

func main() {
	/*
		maxRangeToCheck := 128

		numbersWithSixFactors := getNumbersWithSixFactors(maxRangeToCheck)

		fmt.Printf("numbers with six factors : %v\n", numbersWithSixFactors)

		fmt.Printf("H(%v) : %v \n", maxRangeToCheck, len(numbersWithSixFactors))
	*/

	/*
		(first approach)
		maxRangeToCheck := 262144 // 13208, took 2m 57s

		numbersWithSixFactorsAmount := getNumbersWithSixFactors(maxRangeToCheck)
		fmt.Println(numbersWithSixFactorsAmount)
	*/

	// (second approach)
	maxRangeToCheck := 262144 //13208, took 1m 54s

	numbersWithSixFactorsAmount := getNumbersWithSixFactors(maxRangeToCheck)
	fmt.Println(numbersWithSixFactorsAmount)

	/*
		maxRangeToCheck := 134217728

		numbersWithSixFactorsAmount := getNumbersWithSixFactors(maxRangeToCheck)
		fmt.Println(numbersWithSixFactorsAmount)
	*/

}
