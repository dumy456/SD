package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

func node(id int, in <-chan Msg, out chan<- Msg, leaderChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- Msg{Value: id, Done: false}

	for msg := range in {
		if msg.Done {
			if msg.Value != id {
				out <- msg
			}
			return
		}

		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			leaderChan <- id
			out <- Msg{Value: id, Done: true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	channels := make([]chan Msg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan Msg, n*2)
	}

	leaderChan := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		in := channels[i]
		out := channels[(i+1)%n]
		go node(ids[i], in, out, leaderChan, &wg)
	}

	leader := <-leaderChan
	wg.Wait()

	for _, ch := range channels {
		close(ch)
	}

	return leader
}

func main() {
	ids := []int{14, 42, 5, 89, 23, 71, 3}
	leader := RingFlooding(ids)
	fmt.Printf("Liderul ales în inel este nodul cu ID-ul: %d\n", leader)
}
