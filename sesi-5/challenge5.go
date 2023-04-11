package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go introduce("{coba1, coba2, coba3}")
	introduce("{bisa1, bisa2, bisa2}")

	in := 0
	fmt.Scanln(&in)

}

func introduce(data string) {
	for i := 1; i < 5; i++ {
		fmt.Println(data, i)
		randTime := time.Duration(rand.Intn(240))
		time.Sleep(time.Millisecond * randTime)
	}
}
