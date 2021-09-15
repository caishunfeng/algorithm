package main

import "fmt"

func main() {
	runMap()
	runArray()
}

func runArray() {
	var n = []int{1, 2, 3}

	for i, v := range n {
		fmt.Println(i, v)
		if i == 1 {
			n = append(n, 4)
		}
	}
	fmt.Println(n)
}

func runMap() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	for k, v := range m {
		fmt.Println(k, v)
		if k == "A" {
			delete(m, "A")
		}
		if k == "B" {
			m["D"] = 24
		}
	}
	fmt.Println(len(m))
	fmt.Println(m)
}
