package lib

import (
	filex "github.com/clouby/aoc/utils/file"
)

const SEQUENCE = 14

func trackerQueue(queue []string, element string, positionIndex int) ([]string, int) {
	charExisted := make(map[string]bool)
	catchedIndex := 0
	isRepeated := false

	if len(queue) == SEQUENCE {
		queue = queue[1:]
		queue = append(queue, element)

		for _, key := range queue {
			if _, exist := charExisted[key]; !exist {
				charExisted[key] = true
			} else {
				isRepeated = true
				break
			}
		}

		if !isRepeated {
			catchedIndex = positionIndex + 1
		}

	} else {
		queue = append(queue, element)
	}

	return queue, catchedIndex
}

func DaySixExec(pathname string) int {
	queue := make([]string, 0, SEQUENCE-1)
	marker := 0

	scanner, file := filex.Read(pathname)

	for scanner.Scan() {
		for index, character := range scanner.Text() {
			element := string(character)

			queue, marker = trackerQueue(queue, element, index)

			if marker != 0 {
				break
			}

		}
	}

	defer file.Close()

	return marker
}
