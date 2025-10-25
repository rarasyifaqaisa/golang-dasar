package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Println(s)
	fmt.Println("Len : ", len(s))
	fmt.Println("Cap : ", cap(s))

	s = append(s, 10, 20)
	fmt.Println("After append : ", s)

	s2 := make([]int, 3)
	s3 := copy(s2, s)

	fmt.Println("Copied : ", s2)
	fmt.Println("Jumlah element yang tersalin : ", s3)

	number := []int{1, 2, 3, 4, 5}
	s4 := number[1:4]
	fmt.Println("Slice 4 : ", s4)
}
