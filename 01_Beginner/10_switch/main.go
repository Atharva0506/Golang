package main

import "fmt"

func main() {

	lang := "rust"

	switch lang {
	case "Python":
		fmt.Printf("FastAPI,Flask,Django....")
	case "Javscript", "Typescript":
		fmt.Println("React,NodeJS,Next JS, Angular,Bun...")
	case "Java":
		fmt.Println("I dont like it but Springboot,Hibernet,JSP....")
	case "rust", "go":
		fmt.Println("Learning.....<3")
	default:
		fmt.Println("You're Not a Programmer")

	}

	temp := 20

	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp >= 0 && temp <= 20:
		fmt.Println("Cold")
	case temp >= 21 && temp <= 35:
		fmt.Println("Normal")
	case temp > 35:
		fmt.Println("Hot")
	}

	/*
		- If/Else is slower compared to switch if we have too many conditions.
		- Switch cases act more like a direct lookup and don't blindly evaluate every statement.
		- We get the `fallthrough` feature in switch which is helpful for certain patterns.
	*/
}
