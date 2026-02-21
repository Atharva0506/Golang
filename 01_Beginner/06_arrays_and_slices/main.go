package main

import "fmt"

func main() {

	var carBrand = [3]string{"BMW", "TATA", "Audi"}

	fmt.Println(carBrand)

	var games = make([]string, 1)

	games[0] = "GTA5"

	// games[5] = "BGMI" ===> THIS WILL CAUSE AN ERROR (Index out of bounds)
	games = append(games, "valorant", "COC")

	fmt.Println(games[0:2])

	fmt.Println(len(games))

}
