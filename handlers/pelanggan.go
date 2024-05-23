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

func PelangganMenu() {
	// updateCustomer()
	// viewCustomers()
	// insertCustomers()
	// deleteCustomer()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 17), "Menu Pelanggan", strings.Repeat("=", 17))
		fmt.Println("1. Lihat Semua Pelanggan")
		fmt.Println("2. Tambah Pelanggan")
		fmt.Println("3. Ubah Pelanggan")
		fmt.Println("4. Hapus Pelanggan")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		switch choice {
		case 1:
			viewPelanggan()
		case 2:
			insertPelanggan()
		case 3:
			updatePelanggan()
		case 4:
			deletePelanggan()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func deletePelanggan() {
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

	_, err = db.Exec("DELETE FROM pelanggan WHERE id_pelanggan=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Pelanggan tidak ditemukan.")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("Pelanggan berhasil dihapus.")
}

func updatePelanggan() {
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

	fmt.Print("Nama baru: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("No HP baru: ")
	scanner.Scan()
	noHandphone := scanner.Text()

	if len(nama) == 0 && len(noHandphone) == 0 {
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

		if strings.TrimSpace(nama) != "" {
			queryParts = append(queryParts, fmt.Sprintf("nama_pelanggan=$%d", argID))
			args = append(args, nama)
			argID++
		}
		if strings.TrimSpace(noHandphone) != "" {
			queryParts = append(queryParts, fmt.Sprintf("no_handphone=$%d", argID))
			args = append(args, noHandphone)
			argID++
		}

		args = append(args, id)
		query := fmt.Sprintf("UPDATE pelanggan SET %s WHERE id_pelanggan=$%d", strings.Join(queryParts, ", "), argID)

		_, err := db.Exec(query, args...)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Pelanggan tidak ditemukan.")
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Pelanggan berhasil diubah.")
		}
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data tidak diubah.")
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		updatePelanggan()
	}
}

func insertPelanggan() {
	db := database.ConnectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan ID Pelanggan: ")
	scanner.Scan()
	idPelanggan, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ID Pelanggan harus berupa angka.")
	}

	// Cek apakah ID pelanggan sudah ada di database
	if isExists("pelanggan", "id_pelanggan", idPelanggan) {
		fmt.Println("ID Pelanggan sudah ada. Gunakan ID lain.")
		return
	}

	fmt.Print("Masukkan Nama Pelanggan: ")
	scanner.Scan()
	namaPelanggan := scanner.Text()

	fmt.Print("Masukkan No Handphone: ")
	scanner.Scan()
	noHandphone, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("No Handphone harus berupa angka.")
	}

	if len(namaPelanggan) == 0 {
		fmt.Println("Nama tidak boleh kosong.")
		return
	}

	if noHandphone < 10 {
		fmt.Println("No HP harus minimal 10 digit.")
		return
	}

	sqlStatement := "INSERT INTO pelanggan (id_pelanggan, nama_pelanggan, no_handphone) VALUES ($1, $2, $3)"
	_, err = db.Exec(sqlStatement, idPelanggan, namaPelanggan, noHandphone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pelanggan berhasil ditambahkan.")
}

func viewPelanggan() []models.Pelanggan {
	db := database.ConnectDb()
	defer db.Close()

	sqlStatement := "SELECT id_pelanggan, nama_pelanggan, no_handphone FROM pelanggan;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	pelanggans := scanPelanggan(rows)

	return pelanggans

}

func scanPelanggan(rows *sql.Rows) []models.Pelanggan {
	pelanggans := []models.Pelanggan{}
	var err error

	for rows.Next() {
		pelanggan := models.Pelanggan{}
		err := rows.Scan(&pelanggan.IDPelanggan, &pelanggan.Nama, &pelanggan.NoHandphone)
		if err != nil {
			panic(err)
		}
		pelanggans = append(pelanggans, pelanggan)
		fmt.Printf("ID Pelanggan: %d, Nama: %s, Nomor Handphone: %s\n", pelanggan.IDPelanggan, pelanggan.Nama, pelanggan.NoHandphone)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return pelanggans
}
