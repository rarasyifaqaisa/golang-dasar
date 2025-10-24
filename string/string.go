package main

import "fmt"

func main() {
	// deklarasi string dengan kutip ganda
	name := "Rara"
	message := "Welcome to our aplication"

	// deklarasi string dengan backtict
	paragraf := `Hallo, this is example of multi-line
		  string in go`

	// menampilkan nilai string
	fmt.Println("Name		:", name)
	fmt.Println("Message	:", message)
	fmt.Println("Paragraf	:", paragraf)
}
