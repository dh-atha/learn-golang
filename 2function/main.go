package main

import (
	"errors"
	"fmt"
)

func hello(nama string) (string, error) {
	res := "hello " + nama
	errorNama := errors.New("error concat nama")
	return res, errorNama
}

func sum(nilai ...int) {
	var total int
	for _, value := range nilai {
		total += value
	}
	fmt.Println(total)
}

func main() {
	halo, err := hello("atha")
	fmt.Println(err)
	fmt.Println(halo)

	sum(1, 2, 3, 4)
}
