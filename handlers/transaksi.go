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
	"time"
)

func TransaksiMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(strings.Repeat("=", 17), "Menu Transaksi", strings.Repeat("=", 17))
		fmt.Println("1. Lihat Semua Transaksi")
		fmt.Println("2. Lihat Transaksi Sesuai ID")
		fmt.Println("3. Masukkan Transaksi")
		fmt.Println("4. Lihat Pelanggan")
		fmt.Println("5. Lihat Layanan")
		fmt.Println("6. Kembali Ke Menu Utama")
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
			ViewAllTransaksi()
		case 2:
			ViewTransaksi()
		case 3:
			InsertTransaksi()
		case 4:
			viewPelanggan()
		case 5:
			viewLayanan()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func InsertTransaksi() {

	db := database.ConnectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	// Mulai transaksi
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			// Jika terjadi kesalahan, rollback transaksi
			tx.Rollback()
			log.Fatal(err)
		}
	}()

	// Ambil input dari pengguna
	fmt.Print("Masukkan ID Transaksi: ")
	scanner.Scan()
	idTransaksi, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ID Transaksi harus berupa angka.")
	}

	// Cek apakah ID transaksi sudah ada di database
	if isExists("transaksi", "id_transaksi", idTransaksi) {
		fmt.Println("ID Transaksi sudah ada. Gunakan ID lain.")
		return
	}

	fmt.Print("Masukkan ID Pelanggan: ")
	scanner.Scan()
	idPelanggan, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ID Pelanggan harus berupa angka.")
	}

	// Cek apakah ID pelanggan valid
	if !isExists("pelanggan", "id_pelanggan", idPelanggan) {
		fmt.Println("ID Pelanggan tidak valid.")
		return
	}

	fmt.Print("Masukkan Tanggal Mulai (YYYY-MM-DD): ")
	scanner.Scan()
	tanggalMulai := scanner.Text()

	fmt.Print("Masukkan Tanggal Selesai (YYYY-MM-DD): ")
	scanner.Scan()
	tanggalSelesai := scanner.Text()

	fmt.Print("Masukkan Nama Penerima: ")
	scanner.Scan()
	penerima := scanner.Text()

	// Validasi input
	if tanggalMulai == "" || tanggalSelesai == "" || penerima == "" {
		fmt.Println("Semua data harus diisi.")
		return
	}

	// Format tanggal
	_, err = time.Parse("2006-01-02", tanggalMulai)
	if err != nil {
		fmt.Println("Format tanggal masuk salah. Gunakan format YYYY-MM-DD.")
		return
	}

	_, err = time.Parse("2006-01-02", tanggalSelesai)
	if err != nil {
		fmt.Println("Format tanggal selesai salah. Gunakan format YYYY-MM-DD.")
		return
	}

	// Lakukan operasi INSERT untuk menambahkan data ke tabel transaksi
	sqlStatement := "INSERT INTO transaksi (id_transaksi, id_pelanggan, tanggal_masuk, tanggal_selesai, nama_penerima) VALUES ($1, $2, $3, $4, $5)"
	_, err = tx.Exec(sqlStatement, idTransaksi, idPelanggan, tanggalMulai, tanggalSelesai, penerima)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	// Lakukan operasi INSERT untuk menambahkan data ke tabel detail_transaksi
	for {
		fmt.Print("Masukkan ID Layanan (0 untuk selesai): ")
		scanner.Scan()
		idLayanan, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID Layanan harus berupa angka.")
		}
		if idLayanan == 0 {
			break
		}

		// Cek apakah ID layanan valid
		if !isExists("layanan", "id_layanan", idLayanan) {
			fmt.Println("ID Layanan tidak valid.")
			continue
		}

		fmt.Print("Masukkan Jumlah: ")
		scanner.Scan()
		jumlah, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Jumlah harus berupa angka.")
			continue
		}
		if jumlah <= 0 {
			fmt.Println("Jumlah harus lebih dari 0.")
			continue
		}

		// Lakukan operasi INSERT untuk menambahkan data ke tabel detail_transaksi tanpa id_detail
		sqlDetail := "INSERT INTO detail_transaksi (id_transaksi, id_layanan, jumlah_layanan) VALUES ($1, $2, $3)"
		_, err = tx.Exec(sqlDetail, idTransaksi, idLayanan, jumlah)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	// Jika semua operasi INSERT berhasil, commit transaksi
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaksi berhasil ditambahkan.")
}

func ViewTransaksi() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan ID Transaksi yang ingin dilihat: ")
	scanner.Scan()
	idTransaksiStr := scanner.Text()
	idTransaksi, err := strconv.Atoi(idTransaksiStr)
	if err != nil {
		fmt.Println("ID Transaksi harus berupa angka.")
		return
	}

	transaksis := getTransaksi(idTransaksi)
	printTransaksi(transaksis)
}

func ViewAllTransaksi() {
	transaksis := getTransaksi(0)
	printTransaksi(transaksis)
}

func getTransaksi(idTransaksi int) []models.Transaksi {

	db := database.ConnectDb()
	defer db.Close()

	var rows *sql.Rows
	var err error

	if idTransaksi == 0 {
		sqlStatement := `
		SELECT t.id_transaksi, p.nama_pelanggan, l.layanan, l.satuan, l.harga, dt.jumlah_layanan, t.tanggal_masuk, t.tanggal_selesai, t.nama_penerima,
		       (l.harga * dt.jumlah_layanan) as total
		FROM transaksi t
		JOIN pelanggan p ON t.id_pelanggan = p.id_pelanggan
		JOIN detail_transaksi dt ON t.id_transaksi = dt.id_transaksi
		JOIN layanan l ON dt.id_layanan = l.id_layanan;`
		rows, err = db.Query(sqlStatement)
	} else {
		sqlStatement := `
		SELECT t.id_transaksi, p.nama_pelanggan, l.layanan, l.satuan, l.harga, dt.jumlah_layanan, t.tanggal_masuk, t.tanggal_selesai, t.nama_penerima,
		       (l.harga * dt.jumlah_layanan) as total
		FROM transaksi t
		JOIN pelanggan p ON t.id_pelanggan = p.id_pelanggan
		JOIN detail_transaksi dt ON t.id_transaksi = dt.id_transaksi
		JOIN layanan l ON dt.id_layanan = l.id_layanan
		WHERE t.id_transaksi = $1;`
		rows, err = db.Query(sqlStatement, idTransaksi)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return scanTransaksi(rows)
}

func scanTransaksi(rows *sql.Rows) []models.Transaksi {
	transaksis := []models.Transaksi{}
	var err error

	for rows.Next() {
		transaksi := models.Transaksi{}
		err = rows.Scan(&transaksi.ID, &transaksi.Pelanggan.Nama, &transaksi.Layanan.NamaLayanan, &transaksi.Layanan.Satuan,
			&transaksi.Layanan.HargaSatuan, &transaksi.Jumlah, &transaksi.TanggalMulai, &transaksi.TanggalSelesai,
			&transaksi.Penerima, &transaksi.Total)
		if err != nil {
			log.Fatal(err)
		}
		transaksis = append(transaksis, transaksi)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return transaksis
}

func printTransaksi(transaksis []models.Transaksi) {
	grandTotal := 0

	for _, transaksi := range transaksis {
		fmt.Println(strings.Repeat("=", 150))

		fmt.Printf("ID Transaksi: %d, Nama Pelanggan: %s, Layanan: %s, Satuan: %s, Harga: %d, Jumlah: %d, Tanggal Mulai: %s, Tanggal Selesai: %s, Penerima: %s, Total: %d\n",
			transaksi.ID, transaksi.Pelanggan.Nama, transaksi.Layanan.NamaLayanan, transaksi.Layanan.Satuan,
			transaksi.Layanan.HargaSatuan, transaksi.Jumlah, transaksi.TanggalMulai, transaksi.TanggalSelesai,
			transaksi.Penerima, transaksi.Total)
		grandTotal += transaksi.Total
	}
	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf("Total Transaksi Keseluruhan: %d\n", grandTotal)

}
