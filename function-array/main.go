package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Jumlah elemen : ", len(number))
	fmt.Println("Index 1 : ", number[0])

	number[0] = 80
	fmt.Println("Index 1 : ", number[0])

	fmt.Println("This is array : ")
	for index, value := range number {
		fmt.Println("Value index : ", index, " = ", value)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	fmt.Println("arr1 == arr2 : ", arr1 == arr2)
	fmt.Println("arr1 != arr2 : ", arr1 != arr2)
}
