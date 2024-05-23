package handlers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"submission-godb/database"
	"submission-godb/models"
)

func LayananMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 18), "Menu Layanan", strings.Repeat("=", 18))
		fmt.Println("1. Lihat Semua Layanan")
		fmt.Println("2. Tambah Layanan")
		fmt.Println("3. Ubah Layanan")
		fmt.Println("4. Hapus Layanan")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewLayanan()
		case 2:
			insertLayanan()
		case 3:
			updateLayanan()
		case 4:
			deleteLayanan()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func deleteLayanan() {
	db := database.ConnectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("ID Layanan: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	_, err = db.Exec("DELETE FROM layanan WHERE id_layanan=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Layanan tidak ditemukan.")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("Layanan berhasil dihapus.")
}

func updateLayanan() {
	db := database.ConnectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("ID Pelanggan: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	fmt.Print("Nama Baru: ")
	scanner.Scan()
	namaLayanan := scanner.Text()

	fmt.Print("Satuan Baru: ")
	scanner.Scan()
	satuanLayanan := scanner.Text()

	fmt.Print("Harga Satuan Baru: ")
	scanner.Scan()
	hargaSatuan := scanner.Text()

	if len(namaLayanan) == 0 && len(hargaSatuan) == 0 && len(satuanLayanan) == 0 {
		fmt.Println("Tidak ada perubahan yang dimasukkan.")
		return
	}

	fmt.Print("Apakah Anda yakin ingin mengubah data pelanggan ini (y/n)? ")
	scanner.Scan()
	choice := scanner.Text()

	if strings.ToLower(choice) == "y" {
		var queryParts []string
		var args []interface{}
		argID := 1

		if strings.TrimSpace(namaLayanan) != "" {
			queryParts = append(queryParts, fmt.Sprintf("layanan=$%d", argID))
			args = append(args, namaLayanan)
			argID++
		}
		if strings.TrimSpace(satuanLayanan) != "" {
			queryParts = append(queryParts, fmt.Sprintf("satuan=$%d", argID))
			args = append(args, satuanLayanan)
			argID++
		}
		if strings.TrimSpace(hargaSatuan) != "" {
			queryParts = append(queryParts, fmt.Sprintf("harga=$%d", argID))
			args = append(args, hargaSatuan)
			argID++
		}

		args = append(args, id)
		query := fmt.Sprintf("UPDATE layanan SET %s WHERE id_layanan=$%d", strings.Join(queryParts, ", "), argID)

		_, err := db.Exec(query, args...)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Layanan tidak ditemukan.")
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Layanan berhasil diubah.")
		}
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data tidak diubah.")
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		updatePelanggan()
	}
}

func insertLayanan() {
	db := database.ConnectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ID Layanan: ")
	scanner.Scan()
	idLayanan, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ID Layanan harus berupa angka.")
	}

	// Cek apakah ID layanan sudah ada di database
	if isExists("layanan", "id_layanan", idLayanan) {
		fmt.Println("ID Layanan sudah ada. Gunakan ID lain.")
		return
	}

	fmt.Print("Nama Layanan: ")
	scanner.Scan()
	namaLayanan := scanner.Text()

	fmt.Print("Satuan: ")
	scanner.Scan()
	satuan := scanner.Text()

	fmt.Print("Harga: ")
	scanner.Scan()
	hargaSatuan, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Harga harus berupa angka.")
	}
	// Validasi input
	if namaLayanan == "" || satuan == "" || hargaSatuan == 0 {
		fmt.Println("Semua data harus diisi dengan benar.")
		return
	}

	sqlStatement := "INSERT INTO layanan (id_layanan, layanan, satuan, harga) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(sqlStatement, idLayanan, namaLayanan, satuan, hargaSatuan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Layanan berhasil ditambahkan.")
}

func viewLayanan() []models.Layanan {
	db := database.ConnectDb()
	defer db.Close()

	sqlStatement := "SELECT id_layanan, layanan, satuan, harga FROM layanan;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	layanans := scanLayanan(rows)

	return layanans

}

func scanLayanan(rows *sql.Rows) []models.Layanan {
	layanans := []models.Layanan{}
	var err error

	for rows.Next() {
		layanan := models.Layanan{}
		err := rows.Scan(&layanan.IDLayanan, &layanan.NamaLayanan, &layanan.Satuan, &layanan.HargaSatuan)
		if err != nil {
			panic(err)
		}
		layanans = append(layanans, layanan)
		fmt.Printf("ID Layanan: %d, Layanan: %s, Satuan: %s, Harga : %d \n", layanan.IDLayanan, layanan.NamaLayanan, layanan.Satuan, layanan.HargaSatuan)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return layanans
}
