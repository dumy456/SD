package main

import "fmt"

func rutina(ch chan string) {
	ch <- "Gata"
}
func main() {
	ch := make(chan string)
	go rutina(ch)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	}
}
