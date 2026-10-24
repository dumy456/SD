package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s int
	fmt.Scan(&s)
	var array [10]int
	m := map[int]int{}
	for i := range n {
		fmt.Scan(&array[i])

		if j, ok := m[s-array[i]]; ok {
			fmt.Println(j, i)
		}

		m[array[i]] = i
	}

}
