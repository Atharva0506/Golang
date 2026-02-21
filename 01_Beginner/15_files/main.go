package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("luffy.txt")

	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	// Job Tip: `defer` waits until the ENTIRE function (`main`) is completely finished.
	// This means the file STAYS OPEN while the rest of the code runs below!
	defer file.Close()

	length, err := file.WriteString("First File Created using Golang!!!")
	if err != nil {
		fmt.Println("Error while writting file: ", err)
		return
	}
	fmt.Println("Successfully Wrote: ", length)

	data, err := os.ReadFile("luffy.txt")
	if err != nil {
		fmt.Println("Error while reading file: ", err)
		return
	}

	fileContent := string(data)

	fmt.Println("File Data:\n", fileContent)

	fmt.Println("Deleting File....")

	// ADVANCED GO GOTCHA (Especially on Windows):
	// Because `defer file.Close()` hasn't run yet (we are still inside `main`),
	// the file is technically still OPEN. Windows will block you from deleting an open file!
	// We MUST manually close it right here before trying to delete it.
	file.Close()

	err = os.Remove("luffy.txt")
	if err != nil {
		fmt.Println("Error while deleting file: ", err)
		return
	}
	fmt.Println("File Deleted")
}
