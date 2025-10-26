package main

import "fmt"

func main() {
	fmt.Println("=== Contoh kombinasi for loop di go ===")

	// for standar
	fmt.Println("\nFor standar : ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while style
	fmt.Println("\nWhile style : ")
	j := 2
	for j <= 5 {
		fmt.Printf("%d ", j)
		j += 2
	}
	fmt.Println()

	// infinite loop
	fmt.Println("\nInfinite loop : ")
	k := 5
	for {
		fmt.Printf("%d ", k)
		k--
		if k == 0 {
			break
		}
	}
	fmt.Println()

	// range loop
	fmt.Println("\nRange loop : ")
	fruit := []string{"Apple", "Orange", "Banana"}
	for index, value := range fruit {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}
