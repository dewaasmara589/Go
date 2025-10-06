package main

import (
	"fmt"
)

var g string = "globalVar"

func printLocal() {
	l := "local"
	fmt.Println(l)
	fmt.Println(g)
}

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	// gunakan := jika varibale tidak didefinisikan (seperti var, const)
	i := 1032049348

	// var i int = 1032049348
	s := "Hello, Wold"
	f := 45.06
	b := 5 > 9

	// tipe variable di deklarasikan setelah nama variable
	array := [4]string{"item1", "item2", "item3"}
	slices := []string{"one", "two", "three"}
	m := map[string]string{"letter": "g", "number": "ten", "symbol": "@"}
	fmt.Println(i)
	fmt.Printf(s + " ")
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(array)
	for loopSlices, n := range slices {
		fmt.Printf("index : %d = %q\n", loopSlices, n)
	}
	fmt.Println(m)

	// tidak dapat mengganti value pada variable := dengan sama-sama := harus dengan = saja
	// i := 123
	i = 123
	fmt.Println("new value : ", i)

	// multiple deklar
	j, k, l := "tes", true, 3.14
	fmt.Printf(j)
	fmt.Println(k)
	fmt.Println(l)

	// Global Variable
	printLocal()
	fmt.Println(g)

	// Varible constants
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)

	// Uint menghilangkan negatif number atau start dari 0
	var maxUint32 uint32 = 4294967295
	fmt.Println(maxUint32)
}
