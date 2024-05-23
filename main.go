package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"submission-godb/handlers"
)

func main() {
	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 17), "Enigma Laundry", strings.Repeat("=", 17))
		fmt.Println("1. Menu Pelanggan")
		fmt.Println("2. Menu Layanan")
		fmt.Println("3. Menu Transaksi")
		fmt.Println("4. Keluar")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			handlers.PelangganMenu()
		case 2:
			handlers.LayananMenu()
		case 3:
			handlers.TransaksiMenu()
		case 4:
			fmt.Println("Terima kasih!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
