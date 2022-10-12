package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	radius float64
}

func (l Lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.radius, 2)
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l Lingkaran) keliling() float64 {
	return 2 * math.Pi * l.radius
}

func (p Persegi) keliling() float64 {
	return 4 * p.sisi
}

func main() {
	var l1 Hitung = Lingkaran{radius: 5}
	var p1 Hitung = Persegi{sisi: 5}

	fmt.Println("Luas Lingkaran :", l1.luas())
	fmt.Println("Keliling Lingkaran :", l1.keliling())
	fmt.Println("Luas Persegi :", p1.luas())
	fmt.Println("Keliling Persegi :", p1.keliling())

}
