package main

import "fmt"

// deklarasi interface
type bangunDatar interface {
	luas() int
	keliling() int
}

// Interface tidak di implementasikan secara explicit, tapi secara implicit. maksudnya selama suatu struct mempunyai
// semua method/function yang ada di suatu interface, maka struct tersebut sudah mengimplementasikan interface

type kotak struct {
	sisi int
}

func (k kotak) luas() int {
	return k.sisi * k.sisi
}

func (k kotak) keliling() int {
	return k.sisi * 4
}

type persegiPanjang struct {
	panjang int
	lebar   int
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// func (p persegiPanjang) luas() int {
// 	return p.panjang * p.lebar
// }

func main() {
	var bangun bangunDatar
	bangun = kotak{sisi: 5}
	fmt.Println(bangun.luas())
	fmt.Println(bangun.keliling())

	// error karna persegi panjang tidak mempunyai fungsi luas, uncomment line 35 - 37
	bangun = persegiPanjang{panjang: 2, lebar: 4}
	fmt.Println(bangun.luas())
	fmt.Println(bangun.keliling())
}
