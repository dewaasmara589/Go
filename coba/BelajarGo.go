package main

import (
	// coba dapat dari nama module di go.mod
	"coba/calculation"
	"coba/learning"
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
	// gunakan := jika varibale ada nilainya dan tidak perlu didefinisikan (seperti var, const)
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
	fmt.Println(m, "\n")

	// tidak dapat mengganti value pada variable := dengan sama-sama := harus dengan = saja
	// i := 123
	i = 123
	fmt.Println("new value : ", i, "\n")

	// multiple deklar
	fmt.Print("--- Multiple Declaration ---\n")
	j, k, l := "tes", true, 3.14
	fmt.Printf(j)
	fmt.Println(k)
	fmt.Println(l, "\n")

	// Global Variable
	fmt.Print("--- Global Variables ---\n")
	printLocal()
	fmt.Println(g, "\n")

	// Varible constants
	fmt.Print("--- Constants Variables ---\n")
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes*leapYear, "\n")

	// Uint menghilangkan negatif number atau start dari 0
	var maxUint32 uint32 = 4294967295
	fmt.Print("--- maxUint32 ---\n")
	fmt.Println(maxUint32, "\n")

	// Use Model and learn to
	fmt.Print("--- Functions ---\n")
	subject := learning.Subject{Id: 1, Title: "Belajar Golang"}

	// %s: Nilai string
	// %d: Nilai desimal (integer)
	// %f: Nilai floating point
	// %t: Nilai boolean
	// %c: Nilai karakte
	// tidak dapat digunakan di Println
	fmt.Printf("Id : %d, Title : %s \n", subject.Id, subject.Title)

	// jika beda file namun sama/satu package tetap dapat dipanggil
	// 1. Executable (main)
	// 2. Library (selain main)
	// sentence := TestAja()

	// fmt.Println(sentence) // akan error jika hanya run BelajarGo.go jadi harus run BelajarGo.go KonsepPackage.go

	// Functions
	result := calculation.Add(1, 2)
	fmt.Println("Add : ", result)

	resultMultiple := calculation.Multiple(7, 5)
	fmt.Println("Multiple : ", resultMultiple, "\n")

	// IF dan Switch
	fmt.Print("--- Condition ---\n")

	score := 60
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println("Grade : ", grade)

	number := 5

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	}

	// Perulangan For
	fmt.Print("\n--- Looping For ---\n")
	for i := 0; i <= 5; i++ {
		fmt.Println("Saya sedang belajar GO : ", i)
	}

	title := "Golang the best lengauage"

	for index, letter := range title {
		fmt.Printf("index : %d, letter : %s \n", index, string(letter))
	}

	// ganti dengan _ jika ada variable yang wajib tapi tidak digunakan
	for _, letter := range title {
		fmt.Printf("%s  ", string(letter))
	}

	fmt.Println("\n\n --- QUIZ ---")
	// QUIZ hilangkan index ganjil
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Printf("%s ", string(letter))
		}
	}

	// QUIZ cetak huruf vokal a, i, u, e, o
	fmt.Printf("\n")
	for _, letter := range title {
		// cara 1
		// if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
		// 	fmt.Printf("%s  ", string(letter))
		// }

		// cara 2
		letterString := string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Printf("%s  ", string(letter))
		}
	}

}
