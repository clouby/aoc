package lib

import (
	"strings"
)

const ABC string = "abcdefghijklmnopqrstuvwxyz"

func generatePriority(values string) map[string]int {
	priorities := make(map[string]int)

	upperValues := strings.ToUpper(values)

	for i, v := range values + upperValues {
		priorities[string(v)] = i + 1
	}

	return priorities
}

func DayThreeExec(pathname string) int {

	sum := 0
	scannerElements, file := readFile(pathname)
	priorities := generatePriority(ABC)

	for scannerElements.Scan() {

		text := scannerElements.Text()

		splitPosition := len(text) / 2

		compartmentOne, compartmentTwo := text[:splitPosition], text[splitPosition:]

		indexAppearedText := strings.IndexAny(compartmentOne, compartmentTwo)

		appearedText := string(text[indexAppearedText])

		sum += priorities[appearedText]
	}

	defer file.Close()

	return sum
}
