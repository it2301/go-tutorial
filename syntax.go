package main

import (
	"fmt"
	"sync"
)

func main() {

	// Using packages
	fmt.Println("Hello, World!") // Everything after the dot has to be capitalized

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
	var num11 byte
	var num12 string

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

	// Arrays
	arr := [5]int{1, 2, 2, 3, 4}
	arr[0] = 0
	arr[1] = 1

	// Slices
	slice := []int{1, 2, 3}
	slice = append(slice, 4)       // Appends to a slice at the end
	slice = append(slice, 5, 6, 7) // Appends multiple values to a slice at the end
	slice = slice[1:3]             // Returns a slice from index 1 to 2
	slice = slice[:2]              // Returns a slice from index 0 to 1
	slice = slice[1:]              // Returns a slice from index 1 to end

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

	// Pointers
	pointer := &num
	*pointer = 1

	// Constants
	const constant = 1

	// Unused variables
	_ = 1

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
	print_hello()
	num = add(1, 2)
	num, num2 = add_sub(1, 2)
	num = sum(1, 2, 3, 4, 5)

	// Defer
	defer fmt.Println("World") // Executes after the function gets to the end
	fmt.Println("Hello")

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

	// Disable unused variable warning
	fmt.Println(num, num2, num3, num4, num5, num6, num7, num8, num9, num10, num11, num12, arr, slice, p, m, pointer, constant)
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
