package main

import (
	"fmt"
)

func fill(index int, arr [5]int, number int) [5]int {
	var array [5]int = arr
	array[index] = number
	return array
}

func main() {

	var emptyarray [5]int
	fmt.Println("empty: ", emptyarray)
	fmt.Println("------------------------------")

	var array [5]int

	for {
		fmt.Println("digite o index")

		var index int
   		fmt.Scan(&index)

		fmt.Println("digite o digite o numero")
		var number int 
		fmt.Scan(&number)

		
		array = fill(index, array, number)

		fmt.Println(array)
		fmt.Println("---------------------")
	}
}
