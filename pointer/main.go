package main

import "fmt"

/*
POINTER EXAMPLE IN GOLANG

Goal of this program:
1. Show how a normal variable is stored in memory.
2. Show how to get the memory address of a variable using '&'.
3. Show how a pointer stores that address.
4. Show how to access the value stored at that address using '*'.

Important Concepts:
- Variable
- Memory Address
- Pointer
- Dereferencing

In Go:
&  -> used to get memory address
*  -> used to access value stored at that address
*/

func main() {

	// ------------------------------------------------
	// STEP 1: Create a normal variable
	// ------------------------------------------------

	// num is a normal variable that stores a string value
	// Go will automatically allocate memory in RAM for this variable
	num := "769376393674894"

	// ------------------------------------------------
	// STEP 2: Create a pointer variable
	// ------------------------------------------------

	// &num means:
	// "Give me the memory address of num"
	// ptr will store the memory location where num exists
	ptr := &num

	// ------------------------------------------------
	// STEP 3: Print the actual value
	// ------------------------------------------------

	// This prints the value stored in variable num
	fmt.Println("Value of num:", num)

	// ------------------------------------------------
	// STEP 4: Print memory address
	// ------------------------------------------------

	// ptr contains the memory address of num
	// Example output: 0xc000014220
	fmt.Println("Address of num:", ptr)

	// ------------------------------------------------
	// STEP 5: Dereference the pointer
	// ------------------------------------------------

	// *ptr means:
	// "Go to the memory address stored in ptr
	// and give me the value stored there"
	fmt.Println("Value at address ptr:", *ptr)

}


