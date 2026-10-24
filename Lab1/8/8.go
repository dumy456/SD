package main

import "fmt"

func main() {
	lista := map[string]int{"Ana Popescu": 1, "Mihai Ionescu": 2, "Bogdan Radu": 1, "Elena Dumitrescu": 3, "Cristian Vasile": 2, "Ioana Enache": 1}
	newmap := make(map[int][]string)
	for nume, dep := range lista {

		newmap[dep] = append(newmap[dep], nume)
	}

	fmt.Println(newmap)
}
