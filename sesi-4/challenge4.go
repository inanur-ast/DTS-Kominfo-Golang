package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	args := os.Args
	print(args)
}

func dataStudent() []student {
	students := []student{
		{1, "Muhammad Ridha Anshari", "Makasar", "Freshgraduate", "Pemograman golang banyak diminati oleh perusahan"},
		{2, "Steven Setiawan", "Jakarta", "System Analyst", "Pemograman golang banyak diminati oleh perusahan"},
		{3, "Ansharullah Widiansyah", "Surabaya", "UI/UX Design", "Pemograman golang sangat populer dan diminati oleh perusahan"},
		{4, "Muhammad Kausar", "Yogyakarta", "Freshgraduate", "Pemograman golang sangat populer dan diminati oleh perusahan"},
		{5, "Daru Thobrani Furqon", "Bandung", "Dev Ops", "Pemograman golang sangat populer dan diminati oleh perusahan"},
	}
	return students
}

func print(arguments []string) {
	startIndex, _ := strconv.Atoi(arguments[1])

	for _, value := range dataStudent() {
		if value.absen == startIndex {
			fmt.Println("Nama		:", value.nama)
			fmt.Println("Alamat		:", value.alamat)
			fmt.Println("Pekerjaan	:", value.pekerjaan)
			fmt.Println("Alasan memilih golang :", value.alasan)
			return
		}
	}
	fmt.Println("Siswa Tidak Ditemukan")
}
