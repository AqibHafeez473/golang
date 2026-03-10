package main

import (
	"fmt"
	"unsafe"
)

// Student struct banaya
// Ye ek real-world object ko represent karta hai
// Har student ke paas name, roll number aur marks honge
type Student struct {
	Name   string // student ka name
	RollNo int    // roll number
	Marks  int    // obtained marks
}

func main() {

	// Student struct ka ek variable create kiya
	s1 := Student{
		Name:   "Ali",
		RollNo: 1,
		Marks:  85,
	}

	// Dusra student
	s2 := Student{
		Name:   "Ahmed",
		RollNo: 2,
		Marks:  90,
	}

	// Slice banaya jisme multiple students store kar sakte hain
	// Slice dynamic array jaisa hota hai
	students := []Student{s1, s2}

	// Students ko print karna
	fmt.Println("Student List")
	fmt.Println("------------")

	for i, student := range students {

		// range loop slice ke har element ko iterate karta hai
		fmt.Println("Student Number:", i+1)
		fmt.Println("Name:", student.Name)
		fmt.Println("Roll No:", student.RollNo)
		fmt.Println("Marks:", student.Marks)
		fmt.Println("--------------------")
	}

	// Struct ki memory size check karna
	// unsafe.Sizeof() RAM me struct ka size bytes me batata hai
	size := unsafe.Sizeof(Student{})

	fmt.Println("Memory size of Student struct:", size, "bytes")

	// Example: agar 1000 students ho
	totalStudents := 1000

	totalMemory := int(size) * totalStudents

	fmt.Println("Memory for 1000 students:", totalMemory, "bytes")
}