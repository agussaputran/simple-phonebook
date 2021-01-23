package main

import (
	"fmt"
	"regexp"
)

// Global var
var contact Phone
var contactBook = []Phone{}
var inputMenu, inputSubMenu, inputNameSementara, inputNumberSementara string

// Phone example
type Phone struct {
	name, number, email string
}

func checkName(name string) bool {
	matchString, _ := regexp.MatchString(`^[a-zA-Z]+$`, name)
	return matchString
}

func checkNumber(number string) bool {
	matchString, _ := regexp.MatchString(`^[\d]+$`, number)
	return matchString
}

func checkEmail(email string) bool {
	matchString, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return matchString
}

func main() {

	for {
		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "!" {
			fmt.Println()
			fmt.Println("Anda keluar dari aplikasi")
			fmt.Println()
			break
		} else if inputMenu == "1" {
			fmt.Println()
			fmt.Println("Anda Masuk Pilihan INPUT NOMOR TELEPON")
			fmt.Println()
			addContactName()
			addContactNumber()
			addContactEmail()
			contactBook = append(contactBook, contact)
		} else if inputMenu == "2" {
			fmt.Println()
			fmt.Println("Anda Masuk Pilihan UPDATE NOMOR TELEPON")
			fmt.Println()
			updateContact()
		} else if inputMenu == "3" {
			contactList()
		} else {
			fmt.Println()
			fmt.Println("Masukkan pilihan yang benar")
			fmt.Println()
			continue
		}
	}
}

func menu() {
	fmt.Println("Pilih Menu")
	fmt.Println("1. Input nomor telepon")
	fmt.Println("2. Update nomor telepon")
	fmt.Println("3. List nomor telepon")
	fmt.Println()
	fmt.Println("Masukkan karakter '!' untuk keluar")
	fmt.Print("Masukkan pilihan menu : ? ")
}

func addContactName() {
	for {
		fmt.Println("Masukkan Nama Anda : ")
		fmt.Scanln(&contact.name)

		// if contact.name == "!"{
		// 	menu()
		// }

		if len(contact.name) < 5 || !checkName(contact.name) {
			fmt.Println()
			fmt.Println("FORMAT NAMA HANYA ALPHABET dan TIDAK BOLEH KURANG dari 5")
			fmt.Println()
			continue
		}
		break
	}
}

func addContactNumber() {
	for {
		fmt.Println("Masukkan Nomor Anda : ")
		fmt.Scanln(&contact.number)

		// if contact.number == "!"{
		// 	menu()
		// }

		if len(contact.number) < 10 || !checkNumber(contact.number) {
			fmt.Println()
			fmt.Println("FORMAT NOMOR HANYA ANGKA DENGAN MINIMAL 10 KARAKTER")
			fmt.Println()
			continue
		}
		break
	}
}

func addContactEmail() {
	for {
		fmt.Println("Masukkan Email Anda : ")
		fmt.Scanln(&contact.email)

		// if contact.email == "!"{
		// 	menu()
		// }

		if !checkEmail(contact.email) {
			fmt.Println()
			fmt.Println("EMAIL YANG DIMASUKKAN TIDAK VALID")
			fmt.Println()
			continue
		}
		break
	}
}

func updateContact() {
	for {
		fmt.Print("Masukkan nama : ")
		fmt.Scanln(&inputNameSementara)
		noData := false

		// if inputNameSementara == "!" {
		// 	menu()
		// }

		for i := 0; i < len(contactBook); i++ {
			if contactBook[i].name == inputNameSementara {
				fmt.Print("Masukkan nomor : ")
				fmt.Scanln(&inputNumberSementara)
				contactBook[i].number = inputNumberSementara
				fmt.Println("KONTAK BERHASIL DI UPDATE")
				fmt.Println(contactBook[i])
				noData = false
				break
			} else {
				noData = true
			}
		}

		if noData == true {
			fmt.Println("NAMA TIDAK ADA")
			continue
		}
		break
	}
}

func contactList() {
	fmt.Println()
	fmt.Println(contactBook)
	fmt.Println()
	// fmt.Print("Masukkan '!' Untuk kembali ke menu : ")
	// fmt.Scanln(&inputSubMenu)
	// if inputSubMenu == "!" {
	// 	fmt.Println()
	// 	menu()
	// }
}
