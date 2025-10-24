package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool -> string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke string : ", str)

	// string -> bool
	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke boolean : ", val)
}
