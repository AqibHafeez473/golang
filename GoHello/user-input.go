package main

import "fmt"

func main() {

	var name string
	var age int	
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello %s, you are %d years old and Go is ready! 🚀\n", name, age)
}	