package main

import "fmt"

func main() {
	a := []int{}
	b := make([]int, 1)

	fmt.Println(a)
	fmt.Println(b)

	a = append(a, 1)

	fmt.Println(a)
}