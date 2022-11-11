package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileToArray(filename string) []string {
	readFile, err := os.Open(filename)
	handle("", err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func handle(msg string, err error) {
	if err != nil {
		fmt.Printf("\n%s: %s", msg, err)
	}
}
