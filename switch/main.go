package main

import "fmt"

func main() {
	number := 10

	switch number {
	case 1:
		fmt.Println("Angka satu")
	case 2:
		fmt.Println("Angka dua")
	case 3:
		fmt.Println("Angka tiga")
	default:
		fmt.Println("Angka tidak dikenal")
	}
}
