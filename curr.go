package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/kr/fs"
)

type Match struct {
	line     int
	pos      int
	filename string
}

func main() {
	walker := fs.Walk("/Users/paul/projects")
	// walker := fs.Walk("./")
	query := "path"
	live := true
	count := 0
	// matches := make([]Match, 0)
	for live {
		live = walker.Step()
		count += 1
		stats := walker.Stat()
		curPath := walker.Path()
		if filterWalker(stats) {
			walker.SkipDir()
		}
		if !stats.IsDir() && stats.Mode() == 420 {
			ch_matches := make(chan Match, 1)

			go readLines(curPath, query, ch_matches)

			go func() {
				for mats := range ch_matches {
					println(mats.filename)

				}
			}()

		}
	}

}

func filterWalker(stats os.FileInfo) bool {
	name := stats.Name()
	if name == "node_modules" || name == "target" || name == "dist" {
		return true
	}
	return false
}

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
