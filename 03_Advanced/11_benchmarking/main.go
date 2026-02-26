package main

import "strings"

func ConcatSlow(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += "a"
	}
	return result
}
func ConcatFast(n int) string {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString("a")
	}
	return result.String()
}
func main() {

}
