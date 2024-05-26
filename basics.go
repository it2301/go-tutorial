package main

import (
	"fmt"
	"strings"
	"sync"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func basics() {
	// Here is an example of how to use external packages in Go
	cow_str, err := cowsay.Say(cowsay.Phrase("Hello, World!")) // Here is used the cowsay.Say function that returns a string and an error
	if err != nil {                                            // If there is an error, it will be printed
		fmt.Println(err)
		return
	}

	fmt.Println(cow_str) // The string will be printed

	// Variable declaration and initialization
	var num int = 1
	var num2 = 2 // Automatically detects variable type
	num3 := 3    // Automatically detects variable type

	// Variable declaration
	var num4 int

	// Variable types
	var num5 int  // Can be used int8, int16, int32, int64
	var num6 uint // Can be used uint8, uint16, uint32, uint64
	var num7 float32
	var num8 float64
	var num9 complex64
	var num10 complex128
	var byte1 byte
	var string1 string
	var boolean bool

	// Assigning value to variable
	num = 1
	num += 1
	num -= 1
	num *= 1
	num /= 1
	num %= 1

	// Increment and decrement
	num++
	num--

	// Pointers
	pointer := &num
	*pointer = 1

	// Constants
	const constant = 1 // They never change

	// Unused variables
	_ = 1 // It is used when you don't want to use a variable that a function returns, for example

	// Using arrays
	arr := [5]int{1, 2, 3, 4, 5} // Here is an array of integers
	fmt.Println(arr)             // The array will be printed
	fmt.Println(arr[0])          // The first element of the array will be printed

	start := 1
	end := 4
	fmt.Println(arr[start:end]) // The elements from the start index to the end index - 1 will be printed
	fmt.Println(arr[:end])      // The elements from the start of the array to the end index - 1 will be printed
	fmt.Println(arr[start:])    // The elements from the start index to the end of the array will be printed
	fmt.Println(len(arr))       // The length of the array will be printed

	// Using slices, it is the same as arrays but with the function append to add elements to it
	slice := []int{1, 2, 3, 4, 5} // Here is a slice of integers
	slice = append(slice, 6)      // The append function is used to add an element to the slice
	fmt.Println(slice)            // The slice will be printed

	// Maps
	m := map[string]int{"one": 1, "two": 2}
	m["three"] = 3     // Adds a new key-value pair to a map
	delete(m, "three") // Deletes a key-value pair from a map
	num = m["one"]     // Accesses a value from a map

	// Structs
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}

	// Strings
	str := "Hello, World!"
	fmt.Println(str)
	fmt.Println(str[0])                            // The first character of the string will be printed
	fmt.Println(str[1:5])                          // The characters from the first index to the fifth index - 1 will be printed
	fmt.Println(len(str))                          // The length of the string will be printed
	strings.Contains(str, "Hello")                 // Checks if the string contains a substring and returns a boolean
	strings.Replace(str, "Hello", "Hi", 1)         // Replaces a substring with another substring and returns a new string
	strings.ToLower(str)                           // Converts the string to lowercase and returns a new string
	strings.ToUpper(str)                           // Converts the string to uppercase and returns a new string
	strings.Split(str, ",")                        // Splits the string by a delimiter and returns a slice of strings
	str = fmt.Sprintf("This is a number: %d", num) // Converts a number to a string

	// Conditional statements
	if num == 1 {
		// Do something
	} else if num == 2 {
		// Do something
	} else {
		// Do something
	}

	// Conditional statements with initialization
	if num := 1; num == 1 {
		// Do something
	}

	// Conditional operators
	if num == 1 && num2 == 2 || num3 == 3 {
		// Do something
	}

	// Switch statements
	switch num {
	case 1:
		// Do something
		break
	case 2:
		// Do something
		break
	default:
		// Do something
		break
	}

	// For loops
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
	}

	// For loops with range
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// While-like loops
	for num < 10 {
		num++
	}

	// Infinite loops
	for {
		break
	}

	// Function calls
	print_hello()             // No return value
	num = add(1, 2)           // Single return value
	num, num2 = add_sub(1, 2) // Multiple return values
	num = sum(1, 2, 3, 4, 5)  // Undefined number of arguments

	// Goroutines
	go func() { // Runs in a separate thread and the program doesn't wait for it to finish
		fmt.Println("Hello")
	}()

	// Go routines with wait group
	var wg sync.WaitGroup
	wg.Add(2)   // Adds 2 goroutines to the wait group
	go func() { // Runs in a separate thread and the program doesn't wait for it to finish
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() { // Runs in a separate thread and the program doesn't wait for it to finish
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait() // Waits for all goroutines to finish

	// Disable unused variable error
	fmt.Println(constant, num, num2, num3, num4, num5, num6, num7, num8, num9, num10, byte1, string1, m, p, boolean)
}

// Function without return value
func print_hello() {
	fmt.Println("Hello")
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func add_sub(a int, b int) (int, int) {
	return a + b, a - b
}

// Function with variadic arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func defer_example() {
	defer fmt.Println("World") // Executes after the function gets to the end
	fmt.Println("Hello")
}
