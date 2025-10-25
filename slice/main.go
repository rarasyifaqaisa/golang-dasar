package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	s1 := arr[:] // ini untuk mengambil seluruh element
	fmt.Println("This is s1 : ", s1)

	s2 := arr[:3]
	fmt.Println("This is s2 : ", s2)

	s3 := arr[2:]
	fmt.Println("This is s2 : ", s3)

	s4 := arr[1:4]
	fmt.Println("This is s2 : ", s4)
}
