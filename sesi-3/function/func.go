package main

import "fmt"

func main() {
	studentLists := print("Ariell", "Nanda", "Mailo", "Schannel", "Marco")
	fmt.Printf("%v", studentLists) //print %v digunakan untuk menampilkan DEFAULT FORMAT
}

func print(names ...string) []map[string]string { //function varadic dpt menerima argumen yg tak terbatas jumlahnya
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}
