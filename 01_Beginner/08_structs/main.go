package main

import "fmt"

type Hero struct {
	Name     string // Capital letter => Public (Exported)
	Level    int    // Public
	isActive bool   // Lowercase letter => Private (Unexported - only accessible inside this package)
}

// Job feature: Passing by pointer (*Hero) prevents Go from making a heavy copy
// of the entire struct in memory. This is highly recommended for structs.
func changeLevel(h *Hero, s string, level int) {
	switch s {
	case "+":
		h.Level += level
	case "-":
		h.Level -= level
	default:
		fmt.Println("Not a valid command")
	}
}

func changeStatus(h *Hero) {
	// Pro-tip: You can toggle a boolean much faster by setting it to "NOT itself"
	h.isActive = !h.isActive
}

func main() {
	loki := Hero{
		Name:     "Loki",
		Level:    99,
		isActive: false,
	}

	changeLevel(&loki, "-", 20)
	changeLevel(&loki, "+", 10)

	// loki is currently inactive (false), this will flip him to true
	changeStatus(&loki)

	// Since loki is now active, this will execute
	if loki.isActive {
		fmt.Printf("Active Hero => %+v\n", loki) // %+v prints the struct WITH its field names
	}
}
