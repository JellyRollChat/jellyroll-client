package main

import (
	"bufio"
	"fmt"
	"os"
)

func directoryExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		handle("Could not create directory: ", err)
	}
}

func fileExists(filename string) bool {
	referencedFile, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !referencedFile.IsDir()
}

// createFile Generic file handler
func createFile(filename string) {
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		handle("", err)
		defer file.Close()
	}
}

// writeFile Generic file handler
func writeFile(filename, textToWrite string) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	_, err = file.WriteString(textToWrite)
	err = file.Sync()
	handle("", err)
}

// writeFileBytes Generic file handler
func writeFileBytes(filename string, bytesToWrite []byte) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	_, err = file.Write(bytesToWrite)
	err = file.Sync()
	handle("", err)
}

// readFile Generic file handler
func readFile(filename string) string {
	text, err := os.ReadFile(filename)
	handle("Couldnt read the file: ", err)
	return string(text)
}

func readFileBytes(filename string) []byte {
	text, err := os.ReadFile(filename)
	handle("Couldnt read the file: ", err)
	return text
}

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
