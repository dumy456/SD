package main

import (
	"fmt"
	"math"
	"sync"
)

type Node struct {
	Value       int
	Vot         float64
	Stare       float64
	VotRezultat int
	intrari     []<-chan float64
	iesiri      []chan<- float64
}

func FloodingMedie(n *Node, rounds int, wg *sync.WaitGroup) {
	defer wg.Done()

	n.Stare = n.Vot
	degIn := len(n.intrari)

	send := func(v float64) {
		for _, out := range n.iesiri {
			out <- v
		}
	}

	recvFromAll := func() []float64 {
		vals := make([]float64, degIn)
		for i, in := range n.intrari {
			vals[i] = <-in
		}
		return vals
	}

	for r := 0; r < rounds; r++ {
		send(n.Stare)

		primite := recvFromAll()

		suma := n.Stare
		for _, v := range primite {
			suma += v
		}
		n.Stare = suma / float64(degIn+1)
	}

	if n.Stare > 0.5 {
		n.VotRezultat = 1
	} else {
		n.VotRezultat = 0
	}
}

func main() {
	c12 := make(chan float64, 5)
	c21 := make(chan float64, 5)
	c23 := make(chan float64, 5)
	c32 := make(chan float64, 5)
	c34 := make(chan float64, 5)
	c43 := make(chan float64, 5)
	c41 := make(chan float64, 5)
	c14 := make(chan float64, 5)

	node1 := &Node{Value: 1, Vot: 1.0, intrari: []<-chan float64{c21, c41}, iesiri: []chan<- float64{c12, c14}}
	node2 := &Node{Value: 2, Vot: 0.0, intrari: []<-chan float64{c12, c32}, iesiri: []chan<- float64{c21, c23}}
	node3 := &Node{Value: 3, Vot: 1.0, intrari: []<-chan float64{c23, c43}, iesiri: []chan<- float64{c32, c34}}
	node4 := &Node{Value: 4, Vot: 1.0, intrari: []<-chan float64{c34, c14}, iesiri: []chan<- float64{c43, c41}}

	nodes := []*Node{node1, node2, node3, node4}
	rounds := 30

	var wg sync.WaitGroup
	wg.Add(len(nodes))

	for _, n := range nodes {
		go FloodingMedie(n, rounds, &wg)
	}

	wg.Wait()

	for _, node := range nodes {
		fmt.Printf("Nod: %d | Vot inițial: %.0f | Stare finală (medie): %.4f | Rezultat vot grup: %d\n",
			node.Value, node.Vot, math.Round(node.Stare*10000)/10000, node.VotRezultat)
	}
}
