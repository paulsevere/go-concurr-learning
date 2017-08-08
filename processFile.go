package main

import (
	"bufio"
	"os"
	"strings"
)

func readLines(path string, query string, ch chan Match) {
	file, err := os.Open(path)
	count := 0
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count += 1
		pos := strings.Index(scanner.Text(), query)
		if pos > -1 {
			mat := Match{pos: pos, line: count, filename: path}
			ch <- mat
		}

	}
	close(ch)
	return

}
