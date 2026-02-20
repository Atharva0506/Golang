package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Println("Say my name: ")

	fmt.Scan(&name) // passing addres

	fmt.Println("Your goddam right!! : ", name)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Println("Enter sentence: ")

	sentence, _ := reader.ReadString('\n')
	fmt.Println("You said: ", sentence)

}
