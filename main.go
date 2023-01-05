package main

import (
	"fmt"

	"github.com/clouby/aoc/lib"
)

func main() {
	// Default values for each case
	defaultValueDayFifth := make(map[string][]interface{})

	defaultValueDayFifth["1"] = append(defaultValueDayFifth["1"], "D", "B", "J", "V")
	defaultValueDayFifth["2"] = append(defaultValueDayFifth["2"], "P", "V", "B", "W", "R", "D", "F")
	defaultValueDayFifth["3"] = append(defaultValueDayFifth["3"], "R", "G", "F", "L", "D", "C", "W", "Q")
	defaultValueDayFifth["4"] = append(defaultValueDayFifth["4"], "W", "J", "P", "M", "L", "N", "D", "B")
	defaultValueDayFifth["5"] = append(defaultValueDayFifth["5"], "H", "N", "B", "P", "C", "S", "Q")
	defaultValueDayFifth["6"] = append(defaultValueDayFifth["6"], "R", "D", "B", "S", "N", "G")
	defaultValueDayFifth["7"] = append(defaultValueDayFifth["7"], "Z", "B", "P", "M", "Q", "F", "S", "H")
	defaultValueDayFifth["8"] = append(defaultValueDayFifth["8"], "W", "L", "F")
	defaultValueDayFifth["9"] = append(defaultValueDayFifth["9"], "S", "V", "F", "M", "R")

	// Execute all the challenges
	dayOne := lib.DayOneExec("assets/day1.txt")
	dayTwo := lib.DayTwoExec("assets/day2.txt")
	dayThree := lib.DayThreeExec("assets/day3.txt")
	dayFourth := lib.DayFourthExec("assets/day4.txt")
	dayFifth := lib.DayFifthExec("assets/day5.txt", defaultValueDayFifth)
    daySix := lib.DaySixExec("assets/day6.txt")

	// Print all the results
	fmt.Printf("Day 1: Calorie Counting - Result: %v \n", dayOne)
	fmt.Printf("Day 2: Rock Paper Scissors - Result: %v \n", dayTwo)
	fmt.Printf("Day 3: Rucksack Reorganization - Result: %v \n", dayThree)
	fmt.Printf("Day 4: Camp Cleanup - Result: %v \n", dayFourth)
	fmt.Printf("Day 5: Supply Stacks - Result: %v \n", dayFifth)
    fmt.Printf("Day 6: Tuning Trouble - Result: %v \n", daySix)
}
