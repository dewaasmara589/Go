package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

func main() {
	/*
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

	*/

	/*
		fmt.Println("\n--- Function ---\n")
		SendWelcomeNotification("Dewa")

		cost := CalculateShippingCost(5, 4, 1000)

		fmt.Println(cost)

		var weight float64 = 3.2
		var distance float64 = 3.2
		var ratePerKm float64 = 3.2

		cost = CalculateShippingCost(weight, distance, ratePerKm)

		fmt.Println(cost)

		fmt.Println()
		share, leftover := SplitBill(334000, 3)
		fmt.Println(share)
		fmt.Println(leftover)

		fmt.Println("\n--- Error ---\n")
		_, err := os.Open("test.txt")
		// Test Error File
		// _, err := os.Open("test2.txt")
		if err != nil {
			fmt.Println("Error happened: " + err.Error())
			// return kosong agar program tidak lanjut ke bawah ketika masuk ke if ini
			return
		}

		fmt.Println("File Found")

		fmt.Println()
		// Cara 1 dengan error value
		// result, err := Divide(0, 0)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// fmt.Println(result)

		// Cara 2 dengan error panic
		result2 := Divide2(0, 0)
		fmt.Println("Result: " + fmt.Sprint(result2))

		fmt.Println("\n--- Tipe Data ---\n")
		// cara lain inisialisasi slice
		slcFloat := make([]float64, 10)
		fmt.Println("Slice Size :", fmt.Sprint(len(slcFloat)))

		var studentScore map[string]int = make(map[string]int)
		studentScore = map[string]int{
			"Mario": 90,
			"Budi":  85,
			"Jono":  95,
		}

		// Map akan mengembalikan 2 value : score dan variable boolean untuk cek key
		score, ok := studentScore["Mario123"]
		if ok {
			fmt.Println(score)
		} else {
			fmt.Println("value doesn't exist")
		}

		fmt.Println("\n--- Word Counter Challenge ---\n")
		var word string
		fmt.Printf("Input string: ")
		// fmt.Scanln(&word)

		// scanln tidak menerima value setelah spasi maka gunakan :
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			word = scanner.Text()
		}

		fmt.Println("Your String : " + word)

		var myMap map[string]int = make(map[string]int)
		for _, charachter := range string(word) {
			getValue, cekKey := myMap[string(charachter)]
			if cekKey {
				myMap[string(charachter)] = getValue + 1
			} else {
				myMap[string(charachter)] = 1
			}
		}

		for index, value := range myMap {
			fmt.Printf("%s: %d\n", index, value)
		}

		fmt.Println("\n--- Kunci Jawaban ---")

		strCounter := map[string]int{}
		// rune adalah tipe char atau satu huruf seperti "c"
		// for _ int, c rune := range word{
		for _, c := range word {
			strCounter[string(c)] += 1
		}

		for char, count := range strCounter {
			fmt.Printf("%s: %d\n", char, count)
		}

		fmt.Println("\n--- Pointer ---")

		std := Student{
			Id:   123,
			Name: "Dewa",
			Address: AddressData{
				Id:   23,
				Name: "Jl. Bromo",
			},
		}

		fmt.Println(std)
		fmt.Println(std.Address.Name)
		fmt.Println(std.GetString())

		fmt.Println("\n--- Goroutine ---")
		// Sinkronisasi WaitGroup
		var wg sync.WaitGroup
		wg.Add(1)
		go hello(&wg)
		wg.Wait()
		fmt.Println("Hello from Main!")

	*/

	fmt.Println("\n--- Project 2 ---\n")

	// var pilihanMenu string
	// var loopingMenu bool = true
	// mapContact := map[int]string{}

	// for loopingMenu {
	// 	fmt.Println("\n=====================")
	// 	fmt.Println("INPUT CONTACT")
	// 	fmt.Println("=====================")
	// 	fmt.Println("Menu:")
	// 	fmt.Println("1. Lihat Kontak")
	// 	fmt.Println("2. Tambah Kontak")
	// 	fmt.Println("3. Hapus Kontak")
	// 	fmt.Println("4. Quit")
	// 	fmt.Printf("Pilih menu: ")
	// 	fmt.Scanln(&pilihanMenu)

	// 	if pilihanMenu == "1" {
	// 		fmt.Println()
	// 		if len(mapContact) == 0 {
	// 			fmt.Println("Kontak kosong. Silahkan tambahkan kontak dulu")
	// 		} else {
	// 			for key, value := range mapContact {
	// 				fmt.Printf("%d. %s\n", key, value)
	// 			}
	// 		}
	// 	} else if pilihanMenu == "2" {
	// 		fmt.Println()
	// 		indexContact := 1
	// 		var namaKontak string
	// 		var infoKontak string
	// 		fmt.Printf("Nama Kontak: ")
	// 		fmt.Scanln(&namaKontak)
	// 		fmt.Printf("Informasi Kontak: ")
	// 		fmt.Scanln(&infoKontak)

	// 		mapContact[indexContact+len(mapContact)] = namaKontak + " - " + infoKontak

	// 		fmt.Println("\nTambah kontak sukses")
	// 	} else if pilihanMenu == "3" {
	// 		fmt.Println()
	// 		if len(mapContact) == 0 {
	// 			fmt.Println("Kontak kosong. Silahkan tambahkan kontak dulu")
	// 		} else {
	// 			var idHapusKontak int
	// 			for key, value := range mapContact {
	// 				fmt.Printf("%d. %s\n", key, value)
	// 			}
	// 			fmt.Printf("Input nomor yang hendak dihapus: ")
	// 			fmt.Scan(&idHapusKontak)
	// 			delete(mapContact, idHapusKontak)
	// 		}
	// 	} else if pilihanMenu == "4" {
	// 		fmt.Println()
	// 		fmt.Println("Quitting.. Bye Bye")
	// 		loopingMenu = false
	// 	}
	// }

	// Kunci Jawaban
	var contacts []Contact = make([]Contact, 0)
	for {
		var pilihanMenu int

		fmt.Println("\n=====================")
		fmt.Println("INPUT CONTACT")
		fmt.Println("=====================")
		fmt.Println("Menu:")
		fmt.Println("1. Lihat Kontak")
		fmt.Println("2. Tambah Kontak")
		fmt.Println("3. Hapus Kontak")
		fmt.Println("4. Quit")
		fmt.Printf("Pilih menu: ")
		fmt.Scanln(&pilihanMenu)

		if pilihanMenu == 1 {
			fmt.Println()
			getListKontak(contacts)
		} else if pilihanMenu == 2 {
			var nama, infoKontak string

			fmt.Println()
			fmt.Printf("Nama Kontak: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				nama = scanner.Text()

				fmt.Printf("Informasi Kontak: ")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					infoKontak = scanner.Text()

					newContacts := Contact{Name: nama, ContactInformation: infoKontak}
					contacts = append(contacts, newContacts)
				}
			}
		} else if pilihanMenu == 3 {
			var nomorHapus int

			fmt.Println()
			checkData := getListKontak(contacts)
			if checkData {
				fmt.Printf("Input nomor yang hendak di hapus: ")
				fmt.Scanln(&nomorHapus)

				// default delete numbers[:index]  append  -> numbers[index+1]
				//					 numbers[:3]                  numbers[4:]
				//     		   		[10, 20, 30]                 [50, 90, 60]
				// { 10, 20, 30, 40, 50, 90, 60 } = { 10, 20, 30, 50, 90, 60 }   -> element 40 (index = 3) deleted

				// nomorHapus-1 karena inputan hasil manipulasi string dari index loop yang start dari 0
				contacts = append(contacts[:nomorHapus-1], contacts[nomorHapus:]...)

				fmt.Println("\n Hapus Kontak Sukses")
			}
		} else if pilihanMenu == 4 {
			fmt.Println()
			fmt.Println("Quitting.. Bye Bye")
			os.Exit(0)
		} else {
			fmt.Println()
			fmt.Println("Invalid Input. Bye Bye")
			os.Exit(1)
		}
	}
}

type Contact struct {
	Name               string
	ContactInformation string
}

func getListKontak(contacts []Contact) bool {
	if len(contacts) == 0 {
		fmt.Println("Kontak kosong. Silahkan tambahkan kontak dulu")
		return false
	} else {
		for x := 0; x < len(contacts); x++ {
			fmt.Printf("%d. %s - %s \n", x+1, contacts[x].Name, contacts[x].ContactInformation)
		}
		return true
	}
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from Goroutine!")
}

type AddressData struct {
	Id       int
	Name     string
	City     string
	Province string
}

type Student struct {
	Id      int
	Name    string
	Address AddressData
}

func (s Student) GetString() string {
	return fmt.Sprintf("%d - %s", s.Id, s.Name)
}

func SendWelcomeNotification(username string) {
	fmt.Printf("Hi %s, welcome to the app\n", username)

	// Proses tunggu (1 detik)
	time.Sleep(1 * time.Second)

	fmt.Println("[Logging] Notification successfully sent.")
}

func CalculateShippingCost(weight float64, distance float64, ratePerKm float64) float64 {
	cost := weight * distance * ratePerKm
	return cost
}

func SplitBill(totalAmount float64, numPeople int) (float64, float64) {
	individualSahre := math.Floor(totalAmount/float64(numPeople)*100) / 100

	leftover := totalAmount - (individualSahre * float64(numPeople))

	return individualSahre, leftover
}

func Divide(a, b float64) (float64, error) {
	if b == 0 && a == 0 {
		return 0, errors.New("a and b can't be 0 simultaneously")
	}

	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}

	return a / b, nil
}

func Divide2(a, b float64) float64 {
	if b == 0 && a == 0 {
		panic("a and b can't be 0 simultaneously")
	}

	if b == 0 {
		panic("can't divide by 0")
	}

	return a / b
}
