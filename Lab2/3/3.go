package main

import "fmt"

func rutina(ch chan string) {
	ch <- "Salut"
	ch <- "Ce faci?"
	close(ch)
}
func main() {
	ch := make(chan string, 2)
	go rutina(ch)
	for elem := range ch {
		fmt.Println(elem)
	}
}
