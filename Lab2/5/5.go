package main

import "fmt"

func send(ch chan string) {
	ch <- "Mesaj"
	close(ch)
}

func receive(ch chan string) string {

	return <-ch
}
func main() {
	ch := make(chan string)
	go send(ch)
	go receive(ch)
	fmt.Println(receive(ch))
}
