package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello world"

	// mengubah menjadi huruf kecil
	fmt.Println("Lowercase	: ", strings.ToLower(text))

	// mengubah menjadi huruf besar
	fmt.Println("Uppercase	: ", strings.ToUpper(text))

	// mengecek apakah string dimulai dengan "Hello"
	fmt.Println("Start with hello? ", strings.HasPrefix(text, "Hello"))

	// mengecek apakah mengandung kata "world"
	fmt.Println("Contains world? ", strings.Contains(text, "world"))

	// memisahkan string berdasarkan delimiter
	parts := strings.Split("apple, banana, cerry", ",")
	fmt.Println("Splits		: ", parts)

	// mengganti bagian string
	newtext := strings.ReplaceAll(text, "world", "golang")
	fmt.Println("Replace		: ", newtext)
}
