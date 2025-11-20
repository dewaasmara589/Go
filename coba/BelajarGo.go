package main

import (
	// coba dapat dari nama module di go.mod
	"coba/calculation"
	"coba/learning"
	"errors"
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
	fmt.Printf("1. ")
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Printf("%s ", string(letter))
		}
	}

	// QUIZ cetak huruf vokal a, i, u, e, o
	fmt.Printf("\n2. ")
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

	fmt.Println("\n\n --- Array ---")

	// Cara 1
	// var languages [5] string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Java"
	// languages[3] = "Phyton"
	// languages[4] = "C"

	// Cara 2
	// languages := [5]string{"Go", "Ruby", "Java", "Phyton", "C"}

	// jika vertical harus tambah coma diakhir
	// languages := [5]string{

	// ... default lenght tanpa perlu di deklarasi jumlahnya
	languages := [...]string{
		"Go",
		"Ruby",
		"Java",
		"Phyton",
		"C",
	}

	fmt.Println(languages)
	fmt.Println("length = ", len(languages))

	for index, lang := range languages {
		fmt.Printf("\nindex = %d, language = %s", index, lang)
	}

	fmt.Println("\n\n --- Slice ---")
	// array yang tidak dideklar jumlah valuenya

	// Cara 1 Biasa digunakan karena dinamis
	var gameConsoles []string

	gameConsoles = append(gameConsoles, "PS5")
	gameConsoles = append(gameConsoles, "Xbox")
	gameConsoles = append(gameConsoles, "Nitendo")

	// Cara 2
	// gameConsoles := []string{"PS4", "PS5"}

	fmt.Println(gameConsoles)

	for _, consol := range gameConsoles {
		println(consol)
	}

	fmt.Println("\n\n --- MAP ---")

	// Cara 1
	// key = string
	// value = int
	var myMap map[string]int

	// perlu insialisasi map
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["java"] = 8
	myMap["go"] = 10

	fmt.Println(myMap)
	fmt.Println("value key 'ruby' = ", myMap["ruby"])

	// Cara 2
	myMap2 := map[string]string{"ruby": "is beatuful", "go": "is super fast", "javascript": "hype"}
	fmt.Println(myMap2)

	for key, value := range myMap2 {
		fmt.Println("Key : ", key, ", value : ", value)
	}

	// delete value dengan key di Map
	delete(myMap2, "ruby")
	fmt.Println("Delete key ruby : ", myMap2)

	value := myMap2["javascript"]
	fmt.Println("value key javascript : ", value)

	// cek value dengan isAvailable
	value, isAvailable := myMap2["python"]
	fmt.Println("check value python in myMap2 : ", isAvailable)

	fmt.Println("\n\n --- Slice of MAP ---")
	students := []map[string]string{
		{"name": "test", "score": "A"},
		{"name": "coba", "score": "B"},
		{"name": "kocak", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " (", student["score"], ")")
	}

	fmt.Println("\n\n --- QUIZ ---")
	// Quiz 1 (Hitung Rata-Rata)
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// 	Quiz 2 (masukan ke goodScores jika >= 90)
	var goodScores []int

	var jumlahScored float64
	for _, score := range scores {
		jumlahScored += float64(score)

		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	ratarata := jumlahScored / float64(len(scores))

	fmt.Println("1. ", jumlahScored, "/", len(scores), " = ", ratarata)
	fmt.Println("2. ", goodScores)

	fmt.Println("\n\n --- Multiple Function ---")

	hasilSum, hasilDiff := calculate(10, 5)
	fmt.Printf("Sum = %d, Diff = %d", hasilSum, hasilDiff)

	_, onlyDiff := calculate(20, 8)
	fmt.Printf("\nOnly Difference: %d\n", onlyDiff)

	fmt.Println("\n\n --- QUIZ Function ---")

	angka := []int{10, 5, 8, 9, 7}
	total := sum(angka)

	fmt.Println("1. Hasil SUM : ", total)

	// result, err := quizcalculate(10, 2, "+")
	// result, err := quizcalculate(10, 2, "-")
	// result, err := quizcalculate(10, 2, "*")
	// result, err := quizcalculate(10, 2, "/")
	result, err := quizcalculate(10, 2, "=")

	if err != nil {
		fmt.Println("2. Error: ", err.Error())
		fmt.Println("Terjadi Kesalahan")
	} else {
		fmt.Println("2. Hasil: ", result)
	}

	fmt.Println("\n\n --- Struct ---")
	// mirip seperti Model pada konsep MVC
	user := User{}

	// Cara 1
	user.ID = 1
	user.FirstName = "Dewa"
	user.LastName = "Putra"
	user.Email = "dewaputra@gmail.com"
	user.isActive = true

	// Cara 2 Bisa acak pengisian nilainya
	user2 := User{
		FirstName: "Test",
		Email:     "test@gmail.com",
		LastName:  "Test2",
		isActive:  false,
		ID:        2,
	}

	// Cara 3 harus urut pengisian nilainya
	user3 := User{3, "coba", "coba2", "coba@gmail.com", true}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

	// Struct sebagai parameter
	displayUser1 := displayUser(user)
	fmt.Println(displayUser1)

	// Embedded Struct
	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}
	displayGroup(group)

	// Method
	// func tidak melekat ke siapa", sedangkan untuk method melekat contohnya pada User
	display := user.display()
	fmt.Println(display)
	fmt.Println(user2.display())

	fmt.Println("\n\n --- QUIZ ---")
	// Mengubah func displayGroup ke methoed
	group.displayMethodGroup()
}

func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func sum(value []int) int {
	var hasil int
	for _, angka := range value {
		hasil += angka
	}
	return hasil
}

func quizcalculate(angka1, angka2 int, karakter string) (int, error) {
	var hasil int
	var errorResult error
	if karakter == "+" {
		hasil = angka1 + angka2
	} else if karakter == "-" {
		hasil = angka1 - angka2
	} else if karakter == "*" {
		hasil = angka1 * angka2
	} else if karakter == "/" {
		hasil = angka1 / angka2
	} else {
		errorResult = errors.New("Unknown operation")
	}

	return hasil, errorResult
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	isActive  bool
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Printf("\n")
	fmt.Printf("Member Count : %d", len(group.Users))

	fmt.Println("\nUsers Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func displayUser(user User) string {
	// Sprintf akan mengembalikan string jadi berbeda dangan Printf
	return fmt.Sprintf("Name %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	isAvailable bool
}

func (user User) display() string {
	fmt.Println("Method")
	return fmt.Sprintf("Name %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func (group Group) displayMethodGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Printf("\n")
	fmt.Printf("Member Count : %d", len(group.Users))

	fmt.Println("\nUsers Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
