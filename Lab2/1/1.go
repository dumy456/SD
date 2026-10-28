package main

import "fmt"

func rutina(ch chan string) {
	ch <- "Salut"
}
func main() {
	ch := make(chan string)
	go rutina(ch)
	fmt.Println(<-ch)
}
