package main

import (
	"Basics/utils"
	"fmt"
)

func addNumbers(a, b int) (res int) {
	return a + b
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("This is the initial building block for my Go language learning step.")
	fmt.Println("I am learning Go to make my own APIs to test in my Flutter projects.")
	fmt.Printf("This Printf() function is just like the one we used in C, it is similar to that.\n")

	utils.Greet("Roshan")
	var sum int = addNumbers(3, 4)
	fmt.Println("The sum is ", sum)
}
