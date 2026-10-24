package main

import (
	"fmt"
	"sort"
)

func main() {
	lista := map[string]int{"Andrei": 190, "Maria": 165, "Ion": 182, "Elena": 172, "George": 178, "Ana": 165}
	nume := make([]string, 0)
	for name := range lista {
		nume = append(nume, name)
	}
	sort.Slice(nume, func(i, j int) bool { return lista[nume[i]] > lista[nume[j]] })
	fmt.Println(nume)

}
