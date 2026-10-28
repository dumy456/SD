package main

import (
	"fmt"
	"sync"
)

func rutina(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	go rutina(ch, &wg)
	wg.Add(1)
	ch <- "Salut"
	wg.Wait()

}
