package main

import "fmt"

type Match struct {
	line     int
	pos      int
	filename string
	text     string
}

func main() {
	fmt.Printf("Found %v matches", len(walkAndRead(onlyText)))

}

func waste(interface{}) {

}
