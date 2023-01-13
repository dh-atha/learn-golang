package main

import (
	"fmt"
	privatestruct "learn1/3struct/privateStruct"
)

// golang modifier access
// private adalah ketika suatu variable/fungsi/property/method/interface/struct huruf depannya bukan kapital
// public adalah ketika suatu variable/fungsi/propertymethod/interface/struct huruf depannya kapital

func main() {
	cobaStruct := privatestruct.Contoh{Nama: "atha", Umur: 1}
	// method cobaStruct.getNama tidak bisa digunakan karna termasuk private
	// cobaStruct.getNama
	// method cobaStruct.GetNama tidak bisa digunakan karna termasuk public
	nama := cobaStruct.GetNama()
	fmt.Println(nama)
	cobaStruct.SetNama("dhaiffathin")
	fmt.Println(cobaStruct.GetNama())

	// contoh encapsulation di go
	// tidak bisa mengakses properti hobi dalam struct
	// fmt.Println(cobaStruct.hobi)
	fmt.Println(cobaStruct.GetHobi()) // print kosong karna memang valuenya kosong
	cobaStruct.SetHobi("basket")
	fmt.Println(cobaStruct.GetHobi())
}
