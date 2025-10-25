package main

import "fmt"

func main() {
	mhs := map[string]string{
		"name":  "Rara",
		"class": "A",
	}

	fmt.Println("Name : ", mhs["name"])
	fmt.Println("Class : ", mhs["class"])

	mhs["major"] = "Informatika"
	fmt.Println("Major : ", mhs["major"])
}
