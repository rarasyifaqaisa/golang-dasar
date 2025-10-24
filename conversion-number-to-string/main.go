package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score int = 95
	var text string = strconv.Itoa(score)

	fmt.Println("Nilai ujian	: ", text)
}
