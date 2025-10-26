package main

import "fmt"

func jumlahAngka(number ...int) int {
	total := 0
	for _, v := range number {
		total += v
	}
	return total
}

func main() {
	fmt.Println(jumlahAngka(1, 2, 3))
	fmt.Println(jumlahAngka(1, 2, 3, 5))
	fmt.Println(jumlahAngka())
}
