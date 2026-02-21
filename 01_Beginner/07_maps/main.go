package main

import "fmt"

func main() {
	// 1. Always use make() for maps to avoid panics in production
	students := make(map[string]int)

	students["Luffy"] = 88
	students["zoro"] = 78
	students["Nami"] = 35
	students["Imu"] = 99

	// 2. delete(map, key) is safe even if the key doesn't exist
	delete(students, "Imu")

	// 3. "Comma ok" idiom: Crucial for production to check if a key actually exists
	// Avoids accidentally using a 0 value when the key is just missing
	if val, ok := students["Imu"]; ok {
		fmt.Printf("Value Exists: %v\n", val)
	} else {
		fmt.Printf("Student Imu does not exist in the map\n")
	}

	fmt.Println("\nKey and Values in maps are:")

	// 4. Map iteration order is intentionally randomized by Go!
	for name, marks := range students {
		fmt.Printf("Name: %v || Marks: %v\n", name, marks)
	}

}
