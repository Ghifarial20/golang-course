// Nama	: Ghifari Ahmad L
// Nomor Peserta :  149368582101-594

package main

import (
	"fmt"
	"os"
)

type Friends struct {
	Id, Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	var list = []Friends{
		{Id: "1", Nama: "Ghifari", Alamat: "Jl.Kemenangan", Pekerjaan: "Direktur", Alasan: "menambah Skill"},
		{Id: "2", Nama: "Ahmad", Alamat: "Jl.Mangga", Pekerjaan: "Karyawan", Alasan: "menambah relasi"},
		{Id: "3", Nama: "Haiqal", Alamat: "Jl.Merdeka", Pekerjaan: "Wirausaha", Alasan: "menambah Ilmu"},
		{Id: "4", Nama: "Irfan", Alamat: "Jl.Sesama", Pekerjaan: "Programmer", Alasan: "mencari pengalaman"},
		{Id: "5", Nama: "Daniel", Alamat: "Jl.Masyarakat", Pekerjaan: "Programmer", Alasan: "menambah Skill"},
	}

	idFriends := os.Args[1]
	showFriends(list, idFriends)

}

func showFriends(lists []Friends, id string) {
	for i := range lists {
		if lists[i].Id == id {
			fmt.Printf("Nama : %s \nAlamat : %s \nPekerjaan : %s \nAlasan : %s \n", lists[i].Nama, lists[i].Alamat, lists[i].Pekerjaan, lists[i].Alasan)
		}
	}
}
