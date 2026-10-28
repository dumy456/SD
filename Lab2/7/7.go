package main

import (
	"fmt"
	"sync"
)

type Pozitie struct {
	X int
	Y int
}

func robot1(canal chan<- Pozitie, wag *sync.WaitGroup, pozitie_robot_2 chan Pozitie) {
	defer wag.Done()

	canal <- Pozitie{X: 1, Y: 1}
	canal <- Pozitie{X: 1, Y: 2}
	close(canal)

	for poz := range pozitie_robot_2 {
		fmt.Println("Robot 1 are ", poz)
	}
}

func robot2(canal chan<- Pozitie, wag *sync.WaitGroup, pozitie_robot_1 chan Pozitie) {
	defer wag.Done()

	canal <- Pozitie{X: 1, Y: 5}
	canal <- Pozitie{X: 2, Y: 5}
	close(canal)

	for poz := range pozitie_robot_1 {
		fmt.Println("Robot 2 are ", poz)
	}
}

func main() {
	wg := sync.WaitGroup{}

	ch1 := make(chan Pozitie, 2)
	ch2 := make(chan Pozitie, 2)

	wg.Add(2)

	go robot1(ch1, &wg, ch2)
	go robot2(ch2, &wg, ch1)

	wg.Wait()

	fmt.Println("Sfarsit!")
}
