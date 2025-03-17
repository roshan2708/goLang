package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Number int
	MailId string
}

func main() {
	person1 := Person{
		Name:   "Roshan",
		Age:    21,
		Number: 9876543210,
		MailId: "test@gmail.com",
	}
	fmt.Println(person1)

}
