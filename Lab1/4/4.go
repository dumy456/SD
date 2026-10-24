package main

import "fmt"

func main() {
	lista := []float32{3.14, 2, 5, 7, 1.2, 8.4, 3.89}
	var n int
	fmt.Scan(&n)
	listan := make([]float32, 0)
	listan = append(listan, lista[0:n]...)
	listan = append(listan, lista[n+1:]...)
	fmt.Println(listan)
}
