package main

import "fmt"

func addNumbers(a, b int) (res int) {
	return a + b
}
func main() {
	var sum int = addNumbers(3, 4)
	fmt.Println("The sum is ", sum)
}
