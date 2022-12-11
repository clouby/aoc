package main

import (
	"fmt"

	"github.com/clouby/aoc/lib"
)

func main() {
    dayOne := lib.DayOneExec("assets/day1.txt")   
    dayTwo := lib.DayTwoExec("assets/day2.txt")
    dayThree := lib.DayThreeExec("assets/day3.txt")
    dayFourth := lib.DayFourthExec("assets/day4.txt")

    fmt.Printf("Day 1: Calorie Counting - Result: %v \n", dayOne)
    fmt.Printf("Day 2: Rock Paper Scissors - Result: %v \n", dayTwo)
    fmt.Printf("Day 3: Rucksack Reorganization - Result: %v \n", dayThree)
    fmt.Printf("Day 4: Camp Cleanup - Result: %v \n", dayFourth)
}
