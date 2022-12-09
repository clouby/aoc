package main

import (
	"fmt"

	"github.com/clouby/aoc/lib"
)

func main() {
    dayOne := lib.DayOneExec("assets/day1.txt")   
    dayTwo := lib.DayTwoExec("assets/day2.txt")

    fmt.Printf("Day 1: Calorie Counting - Result: %v \n", dayOne)
    fmt.Printf("Day 2: Rock Paper Scissors - Result: %v \n", dayTwo)
}
