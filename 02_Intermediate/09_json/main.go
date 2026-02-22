package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name    string  `json:"product_name"`
	Price   float64 `json:"cost"`
	InStock bool    `json:"is_available"`
}

func main() {
	// 1. Unmarshal: JSON bytes -> Go Struct

	// Since that didn't match the `json:"is_available"` tag exactly, it silently ignored it!
	fakePayload := []byte(`{"product_name":"Laptop","cost":999.99,"is_available":false}`)

	var product Product
	err := json.Unmarshal(fakePayload, &product)
	if err != nil {
		fmt.Println("Error While reading JSON data: ", err)
		return
	}

	fmt.Printf("Parsed Go Struct: %+v\n", product)
	fmt.Println("Just the Name:", product.Name)

	// 2. Marshal: Go Struct -> JSON bytes
	mouse := Product{Name: "Mouse", Price: 99.9, InStock: true}

	outBytes, err := json.Marshal(mouse)
	if err != nil {
		fmt.Println("Error while parsing data: ", err)
		return
	}
	fmt.Println("JSON bytes output: ", string(outBytes))

}
