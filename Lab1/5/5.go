package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	slices.Reverse(slice)
	fmt.Println(slice)

}
