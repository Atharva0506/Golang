package main

import "fmt"

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c

}
func main() {
	x := 100
	y := 200
	fmt.Printf("Before Swap: x:%d Adrr:%p | y:%d Adrr:%p\n", x, &x, y, &y)
	swap(&x, &y)
	fmt.Printf("After Swap: x:%d Adrr:%p | y:%d Adrr:%p\n", x, &x, y, &y)

}
