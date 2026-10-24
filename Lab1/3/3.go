package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pare := make([]int, 0)
	for _, num := range lista {
		if num%2 == 0 {
			pare = append(pare, num)
		}
	}
	fmt.Println(pare)
}
