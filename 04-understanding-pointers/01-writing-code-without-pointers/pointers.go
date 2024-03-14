package main

import "fmt"

func main() {
	age := 32                  // Regular Variable
	var agePointer *int = &age // Pointer Variable

	fmt.Println("Age is", age)                 // 32
	fmt.Println("Age Pointer is", agePointer)  // 0xc00000a0c8
	fmt.Println("Age Pointer is", *agePointer) //32
	getAdultYears(agePointer)
	fmt.Println("Age Pointer is", agePointer) //50
}

func getAdultYears(age *int) {
	*age = *age + 18
}
