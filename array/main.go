package main

import "fmt"

func main() {
	var number [3]int = [3]int{10, 20, 30}
	fmt.Println(number)
	fmt.Println(number[1])

	number[1] = 80
	fmt.Println(number)
}
