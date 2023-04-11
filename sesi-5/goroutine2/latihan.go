package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		var data = fmt.Sprintf("data %d\t", i)
		wg.Add(1)
		go printFruit(data, &wg)
	}

	wg.Wait()
}

func printFruit(data string, wg *sync.WaitGroup) {
	fmt.Println(data)
	wg.Done()
}
