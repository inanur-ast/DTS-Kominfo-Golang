package main

import "fmt"

func main() {
	value := "selamat malam"

	for _, v := range value {
		fmt.Printf("%c\n", v) //menampilkan karakter selamat malam
	}

	charCount := make(map[string]int)

	for _, v := range value {
		charCount[string(v)] += 1
	}

	fmt.Println(charCount)
}
