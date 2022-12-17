package lib

import (
	"strconv"
	"strings"

	filex "github.com/clouby/aoc/utils/file"
)

func getSections(text string) (string, string, string, string) {
	sectionLeft, sectionRight, _ := strings.Cut(text, ",")

	sectionLeftFirst, sectionLeftLast, _ := strings.Cut(sectionLeft, "-")
	sectionRightFirst, sectionRightLast, _ := strings.Cut(sectionRight, "-")

	return sectionLeftFirst, sectionLeftLast, sectionRightFirst, sectionRightLast
}

func DayFourthExec(pathname string) int {
	pairs := 0
	scannerElements, file := filex.Read(pathname)

	for scannerElements.Scan() {
		text := scannerElements.Text()

		sectionLeftFirst, sectionLeftLast, sectionRightFirst, sectionRightLast := getSections(text)

		first, _ := strconv.Atoi(sectionLeftFirst)
		second, _ := strconv.Atoi(sectionLeftLast)
		third, _ := strconv.Atoi(sectionRightFirst)
		fourth, _ := strconv.Atoi(sectionRightLast)

		if (first <= third && second >= fourth) || (third <= first && fourth >= second) {
			pairs += 1
		}
	}

	defer file.Close()

	return pairs
}
