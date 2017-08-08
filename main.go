package main

import (
	"os"

	"sync"

	"github.com/kr/fs"
)

type Match struct {
	line     int
	pos      int
	filename string
}

func main() {

	var wg sync.WaitGroup
	// walker := fs.Walk("./")
	args := os.Args[1:]
	query := args[1]
	path := args[0]
	walker := fs.Walk(path)

	live := true
	count := 0
	// matches := make([]Match, 0)
	for live {

		live = walker.Step()

		stats := walker.Stat()
		curPath := walker.Path()
		if filterWalker(stats) {
			walker.SkipDir()
		}

		if !stats.IsDir() && stats.Mode() == 420 {
			func() {

				ch_matches := make(chan Match, 1)
				wg.Add(1)
				go readLines(curPath, query, ch_matches)

				go func() {
					defer wg.Done()
					for mat := range ch_matches {
						waste(mat)
						count += 1
						// println(mat.filename, mat.line, mat.pos)

					}
				}()
			}()
		}
	}
	wg.Wait()
	println(count)

}

func filterWalker(stats os.FileInfo) bool {
	name := stats.Name()
	if name == "node_modules" || name == "target" || name == "dist" {
		return true
	}
	return false
}

func waste(interface{}) {

}
