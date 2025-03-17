package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["Roshan"] = 21
	studentsAge["Sudhanshu"] = 21
	studentsAge["Prem"] = 22
	studentsAge["Ishan"] = 21
	studentsAge["Trevor"] = 45
	fmt.Println(studentsAge)
	delete(studentsAge, "Prem")
	fmt.Println(studentsAge)
	age, exists := studentsAge["Roshan"]
	fmt.Println("Roshan age is", age)
	fmt.Println("do he exist", exists)
	person := map[string]int{
		"Roshan": 55,
		"Albert": 22,
		"Cooper": 45,
		"Jarvis": 54,
	}
	for index, value := range person {
		fmt.Printf("key is %s and value is %d\n", index, value)
	}
}
