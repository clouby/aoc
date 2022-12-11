package lib

import (
	"strings"
)

// Rules
// "A" or "X" - Rock
// "B" or "Y" - Paper
// "C" or "Z" - Scissors

type InputBlock int64

type matches struct {
	loss string
	draw string
	win  string
}

const (
	Rock InputBlock = iota + 1
	Paper
	Scissors
)

var rules = map[string]matches{
	"X": {
		loss: "B",
		draw: "A",
		win:  "C",
	},
	"Y": {
		loss: "C",
		draw: "B",
		win:  "A",
	},
	"Z": {
		loss: "A",
		draw: "C",
		win:  "B",
	},
}

func matchRules(point, nextPoint string) int {
	pointPlayer := 0         // Score Counting
	var shape InputBlock = 0 // Default value

	playerRule := rules[point]

	switch point {
	case "A", "X":
		shape = Rock
	case "B", "Y":
		shape = Paper
	case "C", "Z":
		shape = Scissors
	}

	switch nextPoint {
	case playerRule.loss:
		pointPlayer += int(shape) + 0
	case playerRule.win:
		pointPlayer += int(shape) + 6
	case playerRule.draw:
		pointPlayer += int(shape) + 3
	}

	return pointPlayer
}

func getInput(key []string) (playerOneScore, playerTwoScore string) {
	playerOneScore = key[0]
	playerTwoScore = key[1]

	return
}

func DayTwoExec(pathname string) int {
	score := 0

	scanner, file := readFile(pathname)

	for scanner.Scan() {
		element := scanner.Text()

		p1, p2 := getInput(strings.Split(element, " "))

		score += matchRules(p2, p1)
	}

	defer file.Close()

	return score
}
