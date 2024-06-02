package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func files() {

	// Writes a file with the content "Hello, World!" and the permissions 0644 (readable and writable by the owner and readable by others
	// The function returns an error if something goes wrong, so you need to always check it
	write_err := os.WriteFile("file.txt", []byte("Hello, World!"), 0644)
	if write_err != nil {
		panic(write_err)
	}

	// Reads the content of the file "file.txt"
	content, read_err := os.ReadFile("file.txt")
	if read_err != nil {
		panic(read_err)
	}

	// Prints the content of the file
	fmt.Println(string(content))

	// Deletes the file "file.txt"
	delete_err := os.Remove("file.txt")
	if delete_err != nil {
		panic(delete_err)
	}

	// Saving a struct to a json file
	// The struct is defined below
	type Person struct {
		Name string
		Age  int
	}

	// Creating a new person
	per := Person{Name: "John", Age: 30}

	// Saving the person to a json file
	// We will need the encoding/json package
	// The Marshal function returns a byte slice and an error
	json_data, json_err := json.Marshal(per)
	if json_err != nil {
		panic(json_err)
	}

	// Writing the json data to a file
	write_err = os.WriteFile("person.json", json_data, 0644)
	if write_err != nil {
		panic(write_err)
	}

	// Reading the json data from the file
	json_content, read_err := os.ReadFile("person.json")
	if read_err != nil {
		panic(read_err)
	}

	// Creating a new person struct
	var new_per Person

	// Unmarshalling the json data to the new person struct
	// The Unmarshal function takes a byte slice and a pointer to a struct
	// It returns an error
	json_err = json.Unmarshal(json_content, &new_per)
	if json_err != nil {
		panic(json_err)
	}

	// Printing the new person struct
	fmt.Println(new_per)

}
