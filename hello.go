package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

	a := make([]int, 1)

    c, d := vals()

	a = append(a, c)
	a = append(a, d)

    fmt.Println(a)
}