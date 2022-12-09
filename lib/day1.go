package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(pathname string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(pathname)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner, f
}

func DayOneExec(pathname string) int {
	maxCalories := 0
	counter := 0

	scannerElements, file := readFile(pathname)

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
