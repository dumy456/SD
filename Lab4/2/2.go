package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Value    int
	maxValue int
	isLeader bool
	intrari  []<-chan int
	iesiri   []chan<- int
}

func Flooding(n *Node, total int, wg *sync.WaitGroup) {
	defer wg.Done()

	max := n.Value

	inbox := make(chan int, total*len(n.intrari)+10)
	for _, in := range n.intrari {
		go func(ch <-chan int) {
			for v := range ch {
				inbox <- v
			}
		}(in)
	}

	send := func(v int) {
		for _, out := range n.iesiri {
			out <- v
		}
	}

	recv := func() int {
		v := <-inbox
		if v > max {
			max = v
		}
		return v
	}

	mapa := make(map[int]bool)
	send(n.Value)

	vazut := 1
	mapa[n.Value] = true

	for vazut < total {
		v := recv()
		if !mapa[v] {
			mapa[v] = true
			send(v)
			vazut++
		}
	}

	n.maxValue = max
	if n.maxValue == n.Value {
		n.isLeader = true
	}
}

func main() {
	c12 := make(chan int, 10)
	c23 := make(chan int, 10)
	c31 := make(chan int, 10)
	c14 := make(chan int, 10)
	c42 := make(chan int, 10)

	node1 := &Node{Value: 10, maxValue: -1, intrari: []<-chan int{c31}, iesiri: []chan<- int{c12, c14}}
	node2 := &Node{Value: 85, maxValue: -1, intrari: []<-chan int{c12, c42}, iesiri: []chan<- int{c23}}
	node3 := &Node{Value: 30, maxValue: -1, intrari: []<-chan int{c23}, iesiri: []chan<- int{c31}}
	node4 := &Node{Value: 45, maxValue: -1, intrari: []<-chan int{c14}, iesiri: []chan<- int{c42}}

	nodes := []*Node{node1, node2, node3, node4}
	total := len(nodes)

	var wg sync.WaitGroup
	wg.Add(total)

	for _, n := range nodes {
		go Flooding(n, total, &wg)
	}

	wg.Wait()

	for _, node := range nodes {
		fmt.Printf("Nodul %d are max: %d \n", node.Value, node.maxValue)
	}
}
