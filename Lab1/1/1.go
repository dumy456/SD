package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)
	if n < 2 || (n > 2 && n%2 == 0) {
		fmt.Println("Numarul nu este prim")
		return
	} else {
		for d := 3; d*d <= n; d += 2 {
			if n%d == 0 {
				fmt.Println("Numarul nu este prim")
				return
			}
		}
	}
	fmt.Println("Numarul este prim")
}
