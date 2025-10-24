package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "100"
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error message	: ", err.Error())
	} else {
		fmt.Println("Number	: ", number)
	}
}
