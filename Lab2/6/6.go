package main

import (
	"fmt"
)

func robot(id, final int, done chan<- int) {
	poz := 0
	for poz < final {
		poz += 1
	}
	done <- id
}

func main() {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	go robot(1, 10, ch1)
	go robot(2, 10, ch2)

	select {
	case castigator := <-ch1:
		fmt.Printf("Robotul %d a castigat!\n", castigator)
	case castigator := <-ch2:
		fmt.Printf("Robotul %d a castigat!\n", castigator)
	}
}
