package main

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	reader, err := os.Open(path)
	handleErr(err)

	defer reader.Close()

	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
