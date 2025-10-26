package main

import "fmt"

func main() {
	var number1, number2 int

	// input dari pengguna
	fmt.Print("Enter first number : ")
	fmt.Scanln(&number1)
	fmt.Print("Enter second number : ")
	fmt.Scanln(&number2)

	// latihan perbandingan
	fmt.Println("\n=== Result ===")
	fmt.Printf("%d == %d ? %v\n", number1, number2, number1 == number2)
	fmt.Printf("%d != %d ? %v\n", number1, number2, number1 != number2)
	fmt.Printf("%d > %d ? %v\n", number1, number2, number1 > number2)
	fmt.Printf("%d < %d ? %v\n", number1, number2, number1 < number2)
	fmt.Printf("%d >= %d ? %v\n", number1, number2, number1 >= number2)
	fmt.Printf("%d <= %d ? %v\n", number1, number2, number1 <= number2)
}
