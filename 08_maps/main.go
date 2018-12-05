package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["James"] = "james@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "James": "james@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
