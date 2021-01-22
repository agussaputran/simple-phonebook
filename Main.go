package main

import (
	"fmt"
	"regexp"
)

// var checkName = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

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
	var inputMenu, inputSubMenu, inputNameSementara, inputNumberSementara string
	var contact Phone
	var contactBook = []Phone{}

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
			for {

				fmt.Print("Masukkan Nama Anda : ")
				fmt.Scanln(&contact.name)
				if len(contact.name) < 5 || !checkName(contact.name) {
					fmt.Println()
					fmt.Println("nama salah")
					fmt.Println()
					break
				}

				fmt.Print("Masukkan nomor anda : ")
				fmt.Scanln(&contact.number)
				if len(contact.number) < 10 || !checkNumber(contact.number) {
					fmt.Println()
					fmt.Println("number salah")
					fmt.Println()
					break
				}

				fmt.Print("Masukkan email anda : ")
				fmt.Scanln(&contact.email)
				if checkEmail(contact.email) == false {
					fmt.Println()
					fmt.Println("email : ", contact.email, " tidak valid")
					fmt.Println()
					break
				}
				contactBook = append(contactBook, contact)
				fmt.Println()
				fmt.Print("Apakah anda ingin keluar ? jika YA ketikkan '!' jika tidak ketikkan apapun")
				fmt.Scanln(&inputSubMenu)
				if inputSubMenu == "!" {
					fmt.Println()
					break
				} else {
					fmt.Println()
					continue
				}

			}
		} else if inputMenu == "2" {
			for {
				fmt.Println()
				fmt.Println("Anda Masuk Pilihan UPDATE NOMOR TELEPON")
				fmt.Println()
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&inputNameSementara)

				noData := false

				if inputNameSementara == "!" {
					break
				}

				for i := 0; i < len(contactBook); i++ {
					if contactBook[i].name == inputNameSementara {
						fmt.Print("Masukkan nomor : ")
						fmt.Scanln(&inputNumberSementara)
						contactBook[i].number = inputNumberSementara
						noData = false
						break
					} else {
						noData = true
					}
				}

				if noData == true {
					fmt.Println("NAMA TIDAK ADA")
				}

			}
		} else if inputMenu == "3" {
			for {
				fmt.Println()
				fmt.Println("Anda Masuk Pilihan LIST NOMOR TELEPON")
				fmt.Println()
				fmt.Println(contactBook)
				fmt.Print("Masukkan '!' Untuk kembali ke menu : ")
				fmt.Scanln(&inputSubMenu)
				if inputSubMenu == "!" {
					fmt.Println()
					break
				}
			}
		} else {
			fmt.Println()
			fmt.Println("Masukkan pilihan yang benar")
			fmt.Println()
			continue
		}
	}

	// if listContact[0].name == "agus" {
	// 	fmt.Println(listContact[0])
	// } else {
	// 	fmt.Println("No data")
	// }
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

// func addContact(cName, cNumber, cEmail string) {

// }

// func updateContact() {

// }
