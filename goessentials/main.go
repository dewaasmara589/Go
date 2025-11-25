package main

import (
	"fmt"
	"os"
)

func main() {
	// variable int akan otomatis dibaca int64
	fmt.Println("Hello World\n")

	fmt.Println("--- UserInput ---\n")
	var name string
	var umur int
	fmt.Print("Masukkan Namamu : ")
	fmt.Scanln(&name)
	fmt.Println("Halo " + name)
	fmt.Print("Masukkan Umurmu : ")
	fmt.Scanln(&umur)
	fmt.Println("Umurmu sekarang adalah ", umur)
	fmt.Println("Umurmu 5 tahun ke depan adalah ", umur+5)

	fmt.Println("\n--- Looping For ---\n")
	var correctPassword string = "dewa123"
	var input string

	for input != correctPassword {
		fmt.Print("Enter the password: ")
		fmt.Scanln(&input)

		if input != correctPassword {
			fmt.Println("Password Salah")
		}
	}

	fmt.Println("Password Benar")

	fmt.Println("\n--- Switch Case ---\n")
	var choice int

	fmt.Println("Menu :")
	fmt.Println("1. Play")
	fmt.Println("2. Settings")
	fmt.Println("3. Quit")
	fmt.Printf("Input Menu : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Open the game")
	case 2:
		fmt.Println("Open settings")
	case 3:
		os.Exit(0)
	}
	fmt.Println("\n--- Quiz Flow Control ---\n")
	var batas int = 18

	for i := 1; i < batas; i++ {
		fmt.Print("Angka ", i)
		if i%4 == 0 && i%6 == 0 {
			fmt.Print(" WoofMeow")
		} else if i%4 == 0 {
			fmt.Print(" Woof")
		} else if i%6 == 0 {
			fmt.Print(" Meow")
		}
		fmt.Print("\n")
	}

	fmt.Println("")
	var jumlahBintang int = 5
	for i := 0; i < jumlahBintang; i++ {
		for x := 0; x < jumlahBintang; x++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	fmt.Println("")
	var jumlahBintangSegi3 int = 5
	for i := 0; i < jumlahBintangSegi3; i++ {
		for x := 0; x <= i; x++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	fmt.Println("\n--- Project ---\n")
	fmt.Println("=========================")
	fmt.Println("APLIKASI KONVERSI SUHU")
	fmt.Println("=========================")

	var konfirmasiPengulangan string = "y"
	for konfirmasiPengulangan != "n" {
		var suhu float64
		var jenisKonversi float64
		fmt.Printf("Inputkan suhu (Celcius): ")
		fmt.Scanln(&suhu)

		fmt.Println("1. Fahrenheit")
		fmt.Println("2. Reamur")
		fmt.Println("3. Kelvin")
		fmt.Printf("Inputkan nomor suhu konversi: ")
		fmt.Scanln(&jenisKonversi)

		switch jenisKonversi {
		case 1:
			fmt.Print("Suhu Dalam Fahrenheit :", (suhu*(9.0/5.0))+32.0)
		case 2:
			fmt.Print("Suhu Dalam Reamur :", suhu*(4.0/5.0))
		case 3:
			fmt.Print("Suhu Dalam Kelvin :", suhu+273.15)
		default:
			fmt.Print("Inputan Tidak Ditemukan")
		}

		var ulang string = ""
		fmt.Printf("\nApakah anda ingin lanjut ? (y/n) ")
		fmt.Scanln(&ulang)
		konfirmasiPengulangan = ulang
	}
}
