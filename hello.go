package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 10 

	fmt.Println("map:", m)

	v1 := m["k1"]

	fmt.Println(v1)

	v3 := m["k3"]

	fmt.Println(v3)

	fmt.Println(len(m))

	delete(m, "k2")

	fmt.Println(m)

	_, prs := m["k1"]
    fmt.Println("prs:", prs)	

	n := map[int]string{1: "a", 2: "b"}
    fmt.Println(n)	



}	