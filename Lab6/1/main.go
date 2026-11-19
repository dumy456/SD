package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	Alive     bool
}

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	header := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(header[0])
	s, _ := strconv.Atoi(header[1])

	if !scanner.Scan() {
		return
	}
	tokenFields := strings.Fields(scanner.Text())
	tokens := make([]int, n)
	for i := 0; i < n; i++ {
		tokens[i], _ = strconv.Atoi(tokenFields[i])
	}

	adj := make([][]int, n)
	nodes := make([]*Node, n)

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			neighFields := strings.Fields(scanner.Text())
			for _, nf := range neighFields {
				v, _ := strconv.Atoi(nf)
				adj[i] = append(adj[i], v)
			}
		}
		mem := make(Memory)
		mem[Info{ID: i, Value: tokens[i]}] = struct{}{}
		nodes[i] = &Node{
			ID:        i,
			Neighbors: adj[i],
			M:         mem,
			Alive:     true,
		}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex!")
		return
	}

	t := 0
	crashedCount := 0
	maxIter := (s + 1) * diam

	for {
		if crashedCount < s && (n-crashedCount) > 1 {
			if rand.Float64() < 0.5 {
				var aliveNodes []int
				for _, node := range nodes {
					if node.Alive {
						aliveNodes = append(aliveNodes, node.ID)
					}
				}
				victimID := aliveNodes[rand.Intn(len(aliveNodes))]
				nodes[victimID].Alive = false
				crashedCount++
			}
		}

		// Pasul 1 & 2:
		received := make([]Memory, n)
		for i := 0; i < n; i++ {
			received[i] = make(Memory)
		}

		for _, u := range nodes {
			if !u.Alive {
				continue
			}
			for _, vID := range u.Neighbors {
				if nodes[vID].Alive {
					for info := range u.M {
						received[vID][info] = struct{}{}
					}
				}
			}
		}

		// Pasul 3:
		for i, u := range nodes {
			if !u.Alive {
				continue
			}
			for info := range received[i] {
				u.M[info] = struct{}{}
			}
		}

		// Pasul 4:
		if t > maxIter {
			break
		}

		// Pasul 5:
		t++
	}

	var finalMemory []Info
	for _, u := range nodes {
		if u.Alive {
			for info := range u.M {
				finalMemory = append(finalMemory, info)
			}
			break
		}
	}

	sort.Slice(finalMemory, func(i, j int) bool {
		return finalMemory[i].ID < finalMemory[j].ID
	})

	fmt.Printf("Diametrul grafului: %d\n", diam)
	fmt.Printf("Iteratia de oprire: %d\n", t)
	for i, info := range finalMemory {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("(%d, %d)", info.ID, info.Value)
	}
	fmt.Println()
}
