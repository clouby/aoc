package lib

import (
	"fmt"
	"strconv"
	"strings"

	filex "github.com/clouby/aoc/utils/file"
	"github.com/clouby/aoc/utils/stack"
)

func createBoostStack(mapping map[string][]interface{}) map[string]*stack.Stack {
    boostStack := make(map[string]*stack.Stack)

    for key, values := range mapping {
        boostStack[key] = stack.New()
        boostStack[key].Bulk(values...)
    }

    return boostStack
}

func DayFifthExec(pathname string, mapping map[string][]interface{}) string {

    boostStack := createBoostStack(mapping)

	scannerElements, file := filex.Read(pathname)

	for scannerElements.Scan() {
        lineText := scannerElements.Text()

        steps := strings.Split(lineText, " ")

        move, from, to := steps[1], steps[3], steps[5]

        mountMove, _ := strconv.Atoi(move)

        nodesToMove := boostStack[from].Pick(mountMove)

        boostStack[to].Bulk(nodesToMove...)
	}

    var result string

    for i := 0; i < len(boostStack); i++ {
        key := strconv.Itoa(i + 1)
        value := boostStack[key]
        result += fmt.Sprintf("%v", value.Peek()) 
    }

	defer file.Close()

	return result
}
