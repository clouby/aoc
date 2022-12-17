package file

import (
	"bufio"
	"log"
	"os"
)

func Read(pathname string) (*bufio.Scanner, *os.File) {
    file, err := os.Open(pathname)

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return scanner, file
}
