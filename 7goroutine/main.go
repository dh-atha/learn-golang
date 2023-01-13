package main

import (
	"fmt"
	"time"
)

func hello(nama string) string {
	time.Sleep(time.Second)
	return "hello " + nama
}

func main() {
	fmt.Println("---normal---")

	for i := 0; i < 5; i++ {
		greeting := hello(fmt.Sprint(i, ". atha"))
		fmt.Println(greeting)
	}

	fmt.Println("---goroutine---")

	// uncomment waitGroup untuk memblock program agar jangan selesai dulu sebelum semua goroutine selesai diproses
	// wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		go func(i int) {
			// defer wg.Done()
			greeting := hello(fmt.Sprint(i, ". atha"))
			fmt.Println(greeting)
		}(i)
	}
	// wg.Wait()
}
