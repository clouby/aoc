package lib

import (
	"log"
	"strconv"

	filex "github.com/clouby/aoc/utils/file"
)

func DayOneExec(pathname string) int {
	maxCalories := 0
	counter := 0

	scannerElements, file := filex.Read(pathname)

	for scannerElements.Scan() {

		value := scannerElements.Text()

		if value == "" {
			if counter > maxCalories {
				maxCalories = counter
			}
			counter = 0
		} else {
			element, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
				return 0
			}

			counter += element
		}
	}

	defer file.Close()

	return maxCalories
}
