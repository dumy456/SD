package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	slice := make([]string, 0)
	frec := make(map[string]int)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {

		slice = append(slice, scanner.Text())
	}
	for _, cuvant := range slice {
		count, ok := frec[cuvant]
		if !ok {
			frec[cuvant] = 1
		} else {
			frec[cuvant] = count + 1
		}
	}
	fmt.Println(frec)
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
}
