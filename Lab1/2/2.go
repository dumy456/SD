package main

import (
	"fmt"
)

func main() {
	array := [5]int{15, 23, 5, 8, 40}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				aux := array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	fmt.Println(array)
}
